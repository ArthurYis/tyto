/*
 * @Description:
 * @Author: Arthur
 * @Date: 2019-08-10 11:17:14
 * @LastEditTime: 2019-08-10 11:17:33
 * @LastEditors: Arthur
 */
package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJsonFile(fileDir string, data interface{}) interface{} {
	r, err := os.Open(fileDir)
	if err != nil {
		fmt.Println("ReadJsonFile open json file error:" + err.Error())
	}
	decoder := json.NewDecoder(r)
	err = decoder.Decode(data)
	if err != nil {
		fmt.Println("ReadJsonFile json decode error:" + err.Error())
	}
	return data
}
