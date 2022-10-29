package utils

import (
	"crypto/md5"
	"encoding/hex"
)

//获取md5
func GetMd5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}
