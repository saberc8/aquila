package utils

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
)

// UuIdV4 生成一个uuid的数据
func UuIdV4() string {
	uuidV4 := fmt.Sprintf("%s", uuid.NewV4())
	token := strings.Replace(uuidV4, "-", "", -1)
	return token
}

// GenerateToken 定义生成token的方法
func GenerateToken(id int64, data interface{}) (string, error) {
	token := UuIdV4()
	return token, nil
}

// ParseToken 定义解析token的方法
func ParseToken(token string) string {
	// redis中获取token数据
	value := "redis.Get(token)"
	return string(value)
}
