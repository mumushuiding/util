package util

import (
	"math/rand"
	"strings"
)

// RandomNumbers 返回随机数字
// len 为长度
func RandomNumbers(len int) string {
	var buff strings.Builder
	for i := 0; i < len; i++ {
		buff.WriteString(string(rand.Intn(9)))
	}
	return buff.String()
}
