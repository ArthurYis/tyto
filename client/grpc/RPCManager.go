/*
 * @Description:
 * @Author: Arthur
 * @Date: 2019-09-18 15:01:50
 * @LastEditTime: 2019-12-12 09:29:09
 * @LastEditors: Arthur
 */
package grpc

import (
	"athena.com/tyto/bean"
	"athena.com/tyto/client/define"
	logs "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
	"time"
)

const ADDR_GRPC = define.HOST + define.GRPC_PORT
const DEFAULT_IDEL_CONNS = 32
const DEFAULT_HALF_IDEL_TIME = 64
const DEFAULT_USE_CONNS = 4096

var rpc *RPCManager
var rpcOnce sync.Once

//创建一个空容器  设置最大空闲数和最大连接数 设置每个连接的最大空闲时间
//获取连接时 从空闲连接池中获取连接，如果空闲连接池没有连接  判定连接总数是否超出最大连接数  如果超出 等待连接的释放
//TODO 未限制最大使用连接数
type RPCManager struct {
	clientChans  chan *RPCClient
	idelChans    chan *RPCClient
	clientPools  *pools //正在使用使用中的连接集
	usedPools    *pools //已经释放使用的连接集
	idelPools    *pools //可被使用的空闲连接集
	timer        *time.Timer
	doNext       bool
	MaxIdelConns int //最大空闲连接数 默认使用 DEFAULT_IDEL_CONNS
	MaxUseConns  int //最大使用连接数 默认使用 DEFAULT_USE_CONNS
}

func (self *RPCManager) getRPC() *RPCClient {
oneMoreTime:
	client := <-self.clientChans
	self.next()
	if client.state {
		client.idel = false
		client.using = true
		self.clientPools.add(client)
		return client
	} else {
		goto oneMoreTime
	}
}

//填充空闲连接池
func (self *RPCManager) next() {
	if !self.doNext {
		go func(self *RPCManager) {
			self.doNext = true

			if self.idelPools.len() > 0 {
				if client := self.idelPools.get(); client != nil {
					self.clientChans <- client
				}
			}

		oneMoreTime:

			if len(self.clientChans) < 1 {
				client := <-self.idelChans
				if client.state {
					client.idel = true
					self.clientChans <- client
				} else {
					client = nil
				}
				goto oneMoreTime
			}

			if self.idelPools.len() < self.MaxIdelConns {
				client := <-self.idelChans
				if client.state {
					client.idel = true
					self.idelPools.add(client)
				} else {
					client = nil
				}
				goto oneMoreTime
			}
			self.doNext = false
		}(self)
	}
}

func (self *RPCManager) start() {

	//将需要保持的空闲grpc客户端 预先填充好
	for i := 0; i < self.MaxIdelConns; i++ {
		conn, err := grpc.Dial(ADDR_GRPC, grpc.WithInsecure())
		if err != nil {
			logs.Error("grpc.Dial:" + err.Error())
			time.Sleep(5 * time.Second)
			continue
		}
		if len(self.clientChans) < 1 {
			self.clientChans <- newClient(conn)
		} else {
			self.idelPools.add(newClient(conn))
		}
	}

	//监听idelchan 及时补充所需连接
	go func(self *RPCManager) {
		for {
			if self.usedPools.len() > 0 {
				if client := self.usedPools.get(); client != nil {
					if client.state {
						client.idel = true
						self.idelChans <- client
						continue
					} else {
						client = nil
					}
				}
			} else {
				conn, err := grpc.Dial(ADDR_GRPC, grpc.WithInsecure())
				if err != nil {
					logs.Error("grpc.Dial:" + err.Error())
					time.Sleep(5 * time.Second)
					continue
				}
				self.idelChans <- newClient(conn)
			}
		}
	}(self)
	/*
		//自动填充idelChan
		go func(self *RPCManager) {
			for {
				if len(self.clientPools) < self.MaxUseConns-self.MaxIdelConns-1 {
					if len(self.idelPools) < self.MaxIdelConns {
						conn, err := grpc.Dial(ADDR_GRPC, grpc.WithInsecure())
						if err != nil {
							logs.Error("grpc.Dial:" + err.Error())
							time.Sleep(5 * time.Second)
							continue
						}
						self.idelChans <- newClient(conn)
					} else {
						time.Sleep(5 * time.Second)
					}
				} else {
					time.Sleep(5 * time.Second)
				}
			}
		}(self)

		//自动填充idelPool
		go func(self *RPCManager) {
			for {
				select {
				case client := <-self.idelChans:
					if len(self.idelPools) < self.MaxIdelConns {
						if !client.hasClosed {
							self.idelPools[client] = true
						} else {
							client = nil
						}
					} else {
						client.close()
						client = nil
					}
				}

			}
		}(self)

		time.Sleep(1 * time.Second)
		//自动填充clientChan
		go func(self *RPCManager) {
			for {
				if len(self.idelPools) > 0 {
					for client, _ := range self.idelPools {
						fmt.Println("client.state:", client.state)
						self.clientChans <- client
						break
					}
				} else {
					time.Sleep(5 * time.Second)
				}
			}
		}(self)*/
}

func GetRPCManager() *RPCManager {
	rpcOnce.Do(func() {
		if rpc == nil {
			rpc = &RPCManager{
				clientChans:  make(chan *RPCClient, 1),
				idelChans:    make(chan *RPCClient, 1),
				clientPools:  &pools{pools: make(map[*RPCClient]bool), mux: new(sync.Mutex)},
				usedPools:    &pools{pools: make(map[*RPCClient]bool), mux: new(sync.Mutex)},
				idelPools:    &pools{pools: make(map[*RPCClient]bool), mux: new(sync.Mutex)},
				MaxIdelConns: DEFAULT_IDEL_CONNS,
				MaxUseConns:  DEFAULT_USE_CONNS,
			}
			rpc.start()
		}
	})
	return rpc
}

func (self *RPCManager) Recover(client *RPCClient) {
	client.using = false
	self.clientPools.delete(client)
	self.usedPools.add(client)
	// fmt.Println("recover", self.clientPools.len(), self.usedPools.len())
}

type RPCClient struct {
	client   *grpc.ClientConn //rpc客户端
	timer    *time.Timer      //定时任务计时器
	lastTime time.Time        //最后一次使用时间
	using    bool             //是否正在被占用
	idel     bool             //是否被纳入空闲集
	state    bool             //是否可用
}

func newClient(client *grpc.ClientConn) *RPCClient {
	c := &RPCClient{client: client, state: true, lastTime: time.Now().UTC()}
	c.start(0)
	return c
}

func (self *RPCClient) start(duration time.Duration) {
	self.timer = time.AfterFunc(duration, func() {
		//client保活机制 如果client无效  关闭连接
		//如果 client被占用 或者client被收纳到空闲集中 保持心跳 更新操作时间
		if time.Now().UTC().Sub(self.lastTime) < DEFAULT_HALF_IDEL_TIME*2*time.Second {

			if self.idel || self.using {
				client := bean.NewKeepAliveClient(self.client)
				ctx, cancel := context.WithTimeout(context.Background(), TIME_OUT*time.Second)
				defer cancel()
				_, err := client.Ping(ctx, &bean.Ping{Ping: true})
				if err != nil {
					self.lastTime = time.Now().UTC()
				}
			}
		} else {
			if !self.using {
				self.close()
				return
			}
		}
		self.start(DEFAULT_HALF_IDEL_TIME * time.Second)
	})
}

func (self *RPCClient) stop() {
	self.state = false
}

func (self *RPCClient) close() {
	self.state = false
	self.client.Close()
	self.client = nil
	self.timer.Stop()
	self.timer = nil

}

type pools struct {
	pools map[*RPCClient]bool
	mux   *sync.Mutex
}

func (self *pools) add(client *RPCClient) {
	self.mux.Lock()
	defer self.mux.Unlock()
	self.pools[client] = true
}

func (self *pools) get() *RPCClient {
	self.mux.Lock()
	defer self.mux.Unlock()
	for client, _ := range self.pools {
		delete(self.pools, client)
		return client
	}
	return nil
}

func (self *pools) delete(client *RPCClient) {
	self.mux.Lock()
	defer self.mux.Unlock()
	delete(self.pools, client)
}

func (self *pools) len() int {
	return len(self.pools)
}
