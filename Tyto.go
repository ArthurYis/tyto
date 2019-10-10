/*
 * @Description:
 * @Author: Arthur
 * @Date: 2019-08-24 17:28:25
 * @LastEditTime: 2019-10-10 09:07:29
 * @LastEditors: Arthur
 */
package tyto

import (
	"athena.com/tyto/bean"
	"athena.com/tyto/client/grpc"
	// "athena.com/tyto/client/http"
	"athena.com/tyto/utils"
	"time"
)

//链路跟踪接口
type Tyto interface {
	Trace(string) string //开启一个链路的生命周期
	FlushT()             //一个链路的生命周期的终结

	Span(string) string //开启一个链路中节点
	FlushS(string)      //一个节点的终结

	Tag(*bean.Tag) //在跟踪过程中对某个节点进行注解说明
}

type Client interface {
	Trace(*bean.Trace) (int, string)
	FlushT(*bean.Trace) (int, string)
	Span(*bean.Span) (int, string)
	FlushS(*bean.Span) (int, string)
	Tag(*bean.Tag) (int, string)
}

func Default() Tyto {
	node, _ := utils.NewNode(1)
	return &tytor{
		ID: node,
		// client: &http.HttpClient{},
		client: &grpc.GRPCClient{},
		spans:  make(map[string]*bean.Span),
		tags:   make(map[string]*bean.Tag),
	}
}

//TODO 如果记录链路信息失败，需要将失败的信息记录到 tytor中 定时进行尝试重新记录
type tytor struct {
	client Client
	ID     *utils.Node
	trace  *bean.Trace
	spans  map[string]*bean.Span
	tags   map[string]*bean.Tag
}

func (self *tytor) Trace(secondId string) string {
	trace := &bean.Trace{}
	trace.TraceId = self.ID.Generate().String()
	trace.TraceSTime = time.Now().UTC().UnixNano()
	trace.Logging = false
	self.trace = trace
	go self.client.Trace(trace)
	return trace.TraceId
}
func (self *tytor) FlushT() {
	self.trace.TraceETime = time.Now().UTC().UnixNano()
	go self.client.FlushT(self.trace)
}

func (self *tytor) Span(secondId string) string {
	span := &bean.Span{}
	span.SpanId = self.ID.Generate().String()
	span.TraceId = self.trace.TraceId
	span.SpanSTime = time.Now().UTC().UnixNano()
	self.spans[span.SpanId] = span
	go self.client.Span(span)
	return span.SpanId
}
func (self *tytor) FlushS(spanId string) {
	span := self.spans[spanId]
	span.SpanETime = time.Now().UTC().UnixNano()
	delete(self.spans, spanId)
	go self.client.FlushS(span)

}

func (self *tytor) Tag(tag *bean.Tag) {
	go self.client.Tag(tag)
}
