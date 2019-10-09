/*
 * @Description:
 * @Author: Arthur
 * @Date: 2019-09-26 08:53:58
 * @LastEditTime: 2019-10-09 17:25:29
 * @LastEditors: Arthur
 */
package main

import (
	"athena.com/tyto/bean"
	"athena.com/tyto/client"
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
	tyto := client.Default()
	tyto.Trace(secondId)
	defer tyto.FlushT()
	span(tyto)
}

func span(tyto client.Tyto) {
	spanId := tyto.Span(RandString(10))
	defer tyto.FlushS(spanId)
	span1(tyto)
	tyto.Tag(&bean.Tag{
		SpanId: spanId,
		TagId:  RandString(10),
		Desc:   "tage1",
		Time:   time.Now().UTC().UnixNano(),
	})
}

func span1(tyto client.Tyto) {
	spanId := tyto.Span(RandString(10))
	defer tyto.FlushS(spanId)
	tyto.Tag(&bean.Tag{
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
