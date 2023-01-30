package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("sdafgertfsdfgssvfFGHHAASd")
	result := make([]byte, n)

	// 获取随机数 不添加随机种子，确保每一次的随机数都是不重复的
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
