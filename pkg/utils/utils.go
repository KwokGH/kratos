package utils

import (
	"crypto/md5"
	"encoding/hex"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//获取md5
func GetMd5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Encrypt(b []byte) (string, error) {
	return "", nil
}

func NewID() string {
	return primitive.NewObjectID().Hex()
}
