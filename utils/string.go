package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

/**
获取body的data(json)转换为string
*字节数据转string
*/
func GetDataString(req *http.Request) string {
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "{\"code\": 1,\"msg\": \"failed\"}"
	} else {
		return bytes.NewBuffer(result).String()
	}
}
