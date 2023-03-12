package util

import (
	"math/rand"
	"unicode/utf8"
)

// 获得n位由小写字母组成的随机字符串
func RandString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyz")
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// 获取字符串长度
func GetStringLength(name string) int {
	return utf8.RuneCountInString(name)
}
