/*
 * @Description:
 * @Author: Arthur
 * @Date: 2019-09-24 15:31:58
 * @LastEditTime: 2019-09-29 09:57:14
 * @LastEditors: Arthur
 */
package grpc

import (
	"athena.com/tyto/bean"
	errCode "athena.com/tyto/err_code"
	logs "github.com/cihub/seelog"
	"golang.org/x/net/context"
	"time"
)

const (
	TIME_OUT = 60
)

type GRPCClient struct {
}

func (self *GRPCClient) Trace(trace *bean.Trace) (int, string) {
	client := GetRPCManager().getRPC()
	defer GetRPCManager().Recover(client)
	cc := bean.NewTraceHandlerClient(client.client)
	ctx, cancel := context.WithTimeout(context.Background(), TIME_OUT*time.Second)
	defer cancel()
	// fmt.Println("trace:", trace)
	response, err := cc.AddTrace(ctx, trace)
	if err != nil {
		logs.Error("grpc.Trace error:" + err.Error())
		return errCode.GRPC_METHOD_ERR, err.Error()

	}
	logs.Debugf("grpc.Trace response:code = %d,msg = %s", response.Code, response.Message)
	return int(response.Code), response.Message
}
func (self *GRPCClient) FlushT(trace *bean.Trace) (int, string) {
	client := GetRPCManager().getRPC()
	defer GetRPCManager().Recover(client)
	cc := bean.NewTraceHandlerClient(client.client)
	ctx, cancel := context.WithTimeout(context.Background(), TIME_OUT*time.Second)
	defer cancel()
	response, err := cc.FlushTrace(ctx, trace)
	if err != nil {
		logs.Error("grpc.FlushT error:" + err.Error())
		return errCode.GRPC_METHOD_ERR, err.Error()
	}
	logs.Debugf("grpc.FlushT response:code = %d,msg = %s", response.Code, response.Message)
	return int(response.Code), response.Message
}

func (self *GRPCClient) Span(span *bean.Span) (int, string) {
	client := GetRPCManager().getRPC()
	defer GetRPCManager().Recover(client)
	cc := bean.NewSpanHandlerClient(client.client)
	ctx, cancel := context.WithTimeout(context.Background(), TIME_OUT*time.Second)
	defer cancel()
	response, err := cc.AddSpan(ctx, span)
	if err != nil {
		logs.Error("grpc.Span error:" + err.Error())
		return errCode.GRPC_METHOD_ERR, err.Error()
	}
	logs.Debugf("grpc.Span response:code = %d,msg = %s", response.Code, response.Message)
	return int(response.Code), response.Message
}

func (self *GRPCClient) FlushS(span *bean.Span) (int, string) {
	client := GetRPCManager().getRPC()
	defer GetRPCManager().Recover(client)
	cc := bean.NewSpanHandlerClient(client.client)
	ctx, cancel := context.WithTimeout(context.Background(), TIME_OUT*time.Second)
	defer cancel()
	response, err := cc.FlushSpan(ctx, span)
	if err != nil {
		logs.Error("grpc.FlushS error:" + err.Error())
		return errCode.GRPC_METHOD_ERR, err.Error()
	}
	logs.Debugf("grpc.FlushS response:code = %d,msg = %s", response.Code, response.Message)
	return int(response.Code), response.Message
}

func (self *GRPCClient) Tag(tag *bean.Tag) (int, string) {
	client := GetRPCManager().getRPC()
	defer GetRPCManager().Recover(client)
	cc := bean.NewSpanHandlerClient(client.client)
	ctx, cancel := context.WithTimeout(context.Background(), TIME_OUT*time.Second)
	defer cancel()
	response, err := cc.AddTag(ctx, tag)
	if err != nil {
		logs.Error("grpc.tag error:" + err.Error())
		return errCode.GRPC_METHOD_ERR, err.Error()
	}
	logs.Debugf("grpc.tag response:code = %d,msg = %s", response.Code, response.Message)
	return int(response.Code), response.Message
}
