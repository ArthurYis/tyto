/*
 * @Description:
 * @Author: Arthur
 * @Date: 2019-09-26 08:53:58
 * @LastEditTime: 2019-12-12 10:05:17
 * @LastEditors: Arthur
 */
package main

import (
	"athena.com/tyto"
	"athena.com/tyto/bean"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		// trace(RandString(10))
		trace(RandString(12))
		time.Sleep(time.Second)
	}
	select {}
}

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func trace(secondId string) {
	tytor := tyto.Default("platform_001", secondId)
	traceId := tytor.Trace(bean.Trace{UserId: "userId002", UserName: "userName002", FromId: "0", Timeout: 10, Flag: true})
	defer tytor.FlushT()
	span(tytor, "0")
	tytor.Tag(bean.Tag{
		OwnId:    traceId,
		TagId:    RandString(10),
		Desc:     "trace1",
		Times:    time.Now().UTC().UnixNano(),
		Logging:  true,
		LogLevel: "DEBUG",
	})
}

func span(tytor tyto.Tyto, fromId string) {

	spanId := tytor.Span(bean.Span{Timeout: 10, Flag: true, FromId: fromId})
	defer tytor.FlushS(spanId)
	span1(tytor, spanId)
	tytor.Tag(bean.Tag{
		OwnId:   spanId,
		TagId:   RandString(10),
		Desc:    "tage1",
		Times:   time.Now().UTC().UnixNano(),
		Logging: false,
	})
}

func span1(tytor tyto.Tyto, fromId string) {
	spanId := tytor.Span(bean.Span{Timeout: 10, Flag: true, FromId: fromId})
	defer tytor.FlushS(spanId)
	tytor.Tag(bean.Tag{
		OwnId:   spanId,
		TagId:   RandString(10),
		Desc:    "tage2",
		Times:   time.Now().UTC().UnixNano(),
		Logging: false,
	})
	tytor.Log(bean.Log{Content: "test log Content", Level: "INFO"})
}
