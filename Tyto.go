/*
 * @Description:
 * @Author: Arthur
 * @Date: 2019-08-24 17:28:25
 * @LastEditTime: 2019-12-11 10:02:15
 * @LastEditors: Arthur
 */
package tyto

import (
	"athena.com/tyto/bean"
	"athena.com/tyto/client/grpc"
	"athena.com/tyto/utils"
	logs "github.com/cihub/seelog"
	"os"
	"strings"
	"time"
)

var node *utils.Node

//链路跟踪接口
type Tyto interface {
	Trace(bean.Trace) string //开启一个链路的生命周期
	FlushT()

	Span(bean.Span) string //开启一个链路中节点

	FlushS(string)

	Tag(bean.Tag) //在跟踪过程中对某个点进行注解说明
}

type Client interface {
	Trace(*bean.Trace) (int, string)
	Span(*bean.Span) (int, string)
	Tag(*bean.Tag) (int, string)
}

type tytor struct {
	client   Client      //数据处理
	ID       *utils.Node //ID生成器
	platform string      //平台标识
	customId string      //客户自定义请求标识或者数据标识
	trace    *bean.Trace //
	spans    map[string]*bean.Span
}

func Default(platform, customId string) Tyto {
	return &tytor{
		ID: node,
		// client: &http.HttpClient{},
		client:   &grpc.GRPCClient{},
		platform: platform,
		customId: customId,
		spans:    make(map[string]*bean.Span),
	}
}

func (self *tytor) Trace(t bean.Trace) string {
	trace := &bean.Trace{
		Platform: self.platform,
		CustomId: self.customId,
		TraceId:  self.ID.Generate().String(),
		Timeout:  t.Timeout,
		Times:    time.Now().UTC().UnixNano(),
		Operate:  "S",
		Flag:     t.Flag,
		UserId:   t.UserId,
		UserName: t.UserName,
		Logging:  t.Logging,
	}

	if len(strings.TrimSpace(t.FromId)) == 0 {
		trace.FromId = "0"
	} else {
		trace.FromId = t.FromId
	}
	self.trace = trace
	go self.client.Trace(trace)
	return trace.TraceId
}

func (self *tytor) FlushT() {
	logs.Debug("flushT:" + self.trace.TraceId)
	trace := &bean.Trace{
		Platform: self.platform,
		CustomId: self.customId,
		TraceId:  self.trace.TraceId,
		FromId:   self.trace.FromId,
		Timeout:  self.trace.Timeout,
		Times:    time.Now().UTC().UnixNano(),
		Operate:  "E",
		Flag:     self.trace.Flag,
		UserId:   self.trace.UserId,
		UserName: self.trace.UserName,
		Logging:  self.trace.Logging,
	}
	go self.client.Trace(trace)
}

func (self *tytor) Span(s bean.Span) string {
	span := &bean.Span{
		TraceId:  self.trace.TraceId,
		SpanId:   self.ID.Generate().String(),
		Timeout:  s.Timeout,
		Times:    time.Now().UTC().UnixNano(),
		Operate:  "S",
		Flag:     s.Flag,
		Logging:  s.Logging,
		Platform: self.platform,
	}

	if len(strings.TrimSpace(s.FromId)) == 0 {
		span.FromId = "0"
	} else {
		span.FromId = s.FromId
	}
	self.spans[span.SpanId] = span
	go self.client.Span(span)
	return span.SpanId
}

func (self *tytor) FlushS(spanId string) {
	s := self.spans[spanId]
	span := &bean.Span{
		TraceId:  self.trace.TraceId,
		SpanId:   s.SpanId,
		FromId:   s.FromId,
		Timeout:  s.Timeout,
		Times:    time.Now().UTC().UnixNano(),
		Operate:  "E",
		Flag:     s.Flag,
		Logging:  s.Logging,
		Platform: self.platform,
	}
	delete(self.spans, spanId)
	go self.client.Span(span)
}

func (self *tytor) Tag(t bean.Tag) {
	tag := &bean.Tag{
		OwnId:    t.OwnId,
		TagId:    self.ID.Generate().String(),
		Times:    time.Now().UTC().UnixNano(),
		Desc:     t.Desc,
		Logging:  t.Logging,
		Platform: self.platform,
	}

	go self.client.Tag(tag)
}

func init() {
	n, err := utils.NewNode(1)
	if err != nil {
		logs.Error("NewNode.err:" + err.Error())
		os.Exit(1)
	}
	node = n
}
