package utils

import (
	// "encoding/json"
	"strconv"
)

//通用失败JSON格式
func GetErrorJsonData(status int, msg string) string {
	jsons := "{\"status\":" + strconv.Itoa(status) + ",\"message\":\"" + msg + "\"}"
	return jsons
}
