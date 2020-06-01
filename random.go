package util

import (
	"fmt"
	"math/rand"
	"strings"
)

// RandomNumbers 返回随机数字
// len 为长度
func RandomNumbers(len int) string {
	var buff strings.Builder
	for i := 0; i < len; i++ {
		buff.WriteString(fmt.Sprintf("%d", rand.Intn(9)))
	}
	return buff.String()
}
