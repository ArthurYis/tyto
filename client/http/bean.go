/*
 * @Description:
 * @Author: Arthur
 * @Date: 2019-10-09 14:11:26
 * @LastEditTime: 2019-10-09 14:11:51
 * @LastEditors: Arthur
 */
package http

type BaseResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ResponseTime int64  `json:"ResponseTime"`
}

type DataResponse struct {
	*BaseResponse
	Result interface{}
}
