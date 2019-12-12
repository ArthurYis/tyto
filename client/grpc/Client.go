/*
 * @Description:
 * @Author: Arthur
 * @Date: 2019-09-24 15:31:58
 * @LastEditTime: 2019-12-12 09:27:52
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
	response, err := cc.AddTrace(ctx, trace)
	if err != nil {
		logs.Error("grpc.TraceN error:" + err.Error())
		return errCode.GRPC_METHOD_ERR, err.Error()

	}
	// logs.Debugf("grpc.TraceN response:code = %d,msg = %s", response.Code, response.Message)
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
		logs.Error("grpc.SpanN error:" + err.Error())
		return errCode.GRPC_METHOD_ERR, err.Error()
	}
	// logs.Debugf("grpc.SpanN response:code = %d,msg = %s", response.Code, response.Message)
	return int(response.Code), response.Message
}

func (self *GRPCClient) Tag(tag *bean.Tag) (int, string) {
	client := GetRPCManager().getRPC()
	defer GetRPCManager().Recover(client)
	cc := bean.NewTagHandlerClient(client.client)
	ctx, cancel := context.WithTimeout(context.Background(), TIME_OUT*time.Second)
	defer cancel()
	response, err := cc.AddTag(ctx, tag)
	if err != nil {
		logs.Error("grpc.tagN error:" + err.Error())
		return errCode.GRPC_METHOD_ERR, err.Error()
	}
	// logs.Debugf("grpc.tagN response:code = %d,msg = %s", response.Code, response.Message)
	return int(response.Code), response.Message
}

func (self *GRPCClient) Log(log *bean.Log) (int, string) {
	client := GetRPCManager().getRPC()
	defer GetRPCManager().Recover(client)
	cc := bean.NewLogHandlerClient(client.client)
	ctx, cancel := context.WithTimeout(context.Background(), TIME_OUT*time.Second)
	defer cancel()
	response, err := cc.AddLog(ctx, log)
	if err != nil {
		logs.Error("grpc.Log error:" + err.Error())
		return errCode.GRPC_METHOD_ERR, err.Error()
	}
	// logs.Debugf("grpc.Log response:code = %d,msg = %s", response.Code, response.Message)
	return int(response.Code), response.Message
}
