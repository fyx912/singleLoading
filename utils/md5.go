package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

//根据unix时间生产MD5码为token
func GetToken() (int64, string) {
	cruntime := time.Now().Unix()
	h := md5.New()
	h.Write([]byte(strconv.FormatInt(cruntime, 10)))
	return cruntime, hex.EncodeToString(h.Sum(nil))
}
