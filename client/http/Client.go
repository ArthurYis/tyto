/*
 * @Description:
 * @Author: Arthur
 * @Date: 2019-09-24 15:22:55
 * @LastEditTime: 2019-12-11 10:03:43
 * @LastEditors: Arthur
 */
package http

import (
	"athena.com/tyto/bean"
	"athena.com/tyto/client/define"
	"bytes"
	"encoding/json"
	"fmt"
	logs "github.com/cihub/seelog"
	"io/ioutil"
	"net/http"
)

const HTTP_URL = "http://" + define.HOST + define.HTTP_PORT

type HttpClient struct {
}

func (self *HttpClient) Trace(trace *bean.Trace) (int, string) {
	//POST
	bs, err := json.Marshal(trace)
	if err != nil {
		logs.Error("HttpClient.json.Marshal.err:" + err.Error())
		return -1, ""
	}
	response := do(http.MethodPost, "/tyto/hedwig/trace/add", bs)
	if response != nil {
		return response.Code, response.Message
	}
	return -1, ""
}

func (self *HttpClient) Span(span *bean.Span) (int, string) {
	//POST
	bs, err := json.Marshal(span)
	if err != nil {
		logs.Error("HttpClient.json.Marshal.err:" + err.Error())
		return -1, ""
	}
	response := do(http.MethodPost, "/tyto/hedwig/span/add", bs)
	if response != nil {
		return response.Code, response.Message
	}
	return -1, ""
}

func (self *HttpClient) Tag(tag *bean.Tag) (int, string) {
	//POST
	bs, err := json.Marshal(tag)
	if err != nil {
		logs.Error("HttpClient.json.Marshal.err:" + err.Error())
		return -1, ""
	}
	response := do(http.MethodPost, "/tyto/hedwig/span/tag/add", bs)
	if response != nil {
		return response.Code, response.Message
	}
	return -1, ""
}

func do(method, path string, data []byte) *BaseResponse {
	request, err := http.NewRequest(method, HTTP_URL+path, bytes.NewBuffer(data))
	if err != nil {
		logs.Errorf("client.%s.err: %s", method, err.Error())
		return nil
	}
	request.Header.Set("Content-Type", "application/json;charset=utf-8")

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		logs.Error("client.do.err:" + err.Error())
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("resp.Body.read.err:" + err.Error())
		return nil
	}

	result := new(BaseResponse)
	fmt.Println(string(body))
	err = json.Unmarshal(body, result)
	if err != nil {
		logs.Error("client.Post.response.json.err:" + err.Error())
		return nil
	}
	logs.Debug("result:", result)
	return result
}
