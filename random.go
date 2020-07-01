package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// RandomNumbers 返回随机数字
// len 为长度
func RandomNumbers(len int) string {
	var buff strings.Builder
	r := rand.New(rand.NewSource(time.Now().Local().UnixNano()))
	for i := 0; i < len; i++ {

		buff.WriteString(fmt.Sprintf("%d", r.Intn(9)))
	}
	return buff.String()
}
