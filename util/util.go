package util

import (
	"math/rand"
	"time"
)

// RandString 生成长度为 n 的随机字符串
func RandString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	// 使用 make 创建一个长度为 n 的字节类型的切片
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
