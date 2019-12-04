/*
 * @Description:
 * @Author: Arthur
 * @Date: 2019-09-26 08:53:58
 * @LastEditTime: 2019-12-04 14:41:12
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
	for i := 0; i < 1000; i++ {
		trace(RandString(10))
		time.Sleep(1000 * time.Millisecond)
	}

}

func trace(secondId string) {
	tyto := tyto.Default()
	tyto.Trace(bean.Trace{Platform: "platform_001", SecondId: secondId, UserId: "userId002", UserName: "userName002", Logging: false})
	defer tyto.FlushT()
	span(tyto)
}

func span(tyto tyto.Tyto) {
	spanId := tyto.Span(RandString(10))
	defer tyto.FlushS(spanId)
	span1(tyto)
	tyto.Tag(bean.Tag{
		SpanId: spanId,
		TagId:  RandString(10),
		Desc:   "tage1",
		Time:   time.Now().UTC().UnixNano(),
	})
}

func span1(tyto tyto.Tyto) {
	spanId := tyto.Span(RandString(10))
	defer tyto.FlushS(spanId)
	tyto.Tag(bean.Tag{
		SpanId: spanId,
		TagId:  RandString(10),
		Desc:   "tage2",
		Time:   time.Now().UTC().UnixNano(),
	})
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
