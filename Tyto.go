/*
 * @Description:
 * @Author: Arthur
 * @Date: 2019-08-24 17:28:25
 * @LastEditTime: 2019-12-04 17:19:26
 * @LastEditors: Arthur
 */
package tyto

import (
	"athena.com/tyto/bean"
	"athena.com/tyto/client/grpc"
	"athena.com/tyto/utils"
	"strings"
	"time"
)

//链路跟踪接口
type Tyto interface {
	Trace(bean.Trace) string //开启一个链路的生命周期
	FlushT()                 //一个链路的生命周期的终结

	Span(string) string //开启一个链路中节点
	FlushS(string)      //一个节点的终结

	Tag(bean.Tag) //在跟踪过程中对某个节点进行注解说明
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

func (self *tytor) Trace(t bean.Trace) string {
	trace := &bean.Trace{Subs: &bean.Subs{Spans: make(map[string]bool), Tags: make(map[string]bool)}}
	trace.SecondId = t.SecondId
	trace.UserId = t.UserId
	trace.UserName = t.UserName
	trace.Platform = t.Platform
	trace.TraceId = self.ID.Generate().String()
	trace.TraceSTime = time.Now().UTC().UnixNano()
	trace.Logging = t.Logging
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
	self.trace.TraceETime = time.Now().UTC().UnixNano()
	go self.client.FlushT(self.trace)
}

func (self *tytor) Span(fromId string) string {
	span := &bean.Span{}
	span.SpanId = self.ID.Generate().String()
	span.FromId = fromId
	span.TraceId = self.trace.TraceId
	span.SpanSTime = time.Now().UTC().UnixNano()
	self.spans[span.SpanId] = span
	self.trace.Subs.Spans[span.SpanId] = true
	go self.client.Span(span)
	return span.SpanId
}
func (self *tytor) FlushS(spanId string) {
	span := self.spans[spanId]
	span.SpanETime = time.Now().UTC().UnixNano()
	delete(self.spans, spanId)
	go self.client.FlushS(span)
}

func (self *tytor) Tag(tag bean.Tag) {
	newTag := &bean.Tag{
		TraceId: self.trace.TraceId,
		SpanId:  tag.SpanId,
		TagId:   self.ID.Generate().String(),
		Desc:    tag.Desc,
		Time:    time.Now().UTC().UnixNano(),
	}

	self.trace.Subs.Tags[newTag.TagId] = true
	go self.client.Tag(newTag)
}
