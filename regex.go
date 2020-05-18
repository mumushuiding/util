package util

import (
	"regexp"
	"unicode"
)

// IsMobile 是否是电话号码
func IsMobile(value string) (bool, error) {
	return regexp.MatchString(`^\d{11}$`, value)
}

// IsEmail 是否是邮箱
func IsEmail(value string) (bool, error) {
	return regexp.MatchString(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`, value)
}

// IsChinese 是否是中文
func IsChinese(value string) bool {
	for _, v := range value {
		if !unicode.Is(unicode.Han, v) {
			return false
		}
	}
	return true
}

// IsABC 是否是英文字母
func IsABC(value string) (bool, error) {
	return regexp.MatchString(`[\w]$`, value)
}
