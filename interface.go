package util

import (
	"errors"
	"fmt"
	"reflect"
)

// Interface2Int interface{} 转换成数字
func Interface2Int(value interface{}) (int, error) {
	if value == nil {
		return 0, errors.New("interface转换成数字失败:interface 值为nil")
	}
	switch value.(type) {
	case int:
		return value.(int), nil
	case float64:
		return int(value.(float64)), nil
	case float32:
		return int(value.(float32)), nil
	default:
		return 0, errors.New("interface转换成数字失败:不是数字类型")
	}
}

// Interface2String 对象转字符串
func Interface2String(value interface{}) (string, error) {
	if value == nil {
		return "", errors.New("interface转换成字符串失败:interface 值为nil")
	}
	switch value.(type) {
	case string:
		if len(value.(string)) == 0 {
			return "", errors.New("字符串不能为空")
		}
		return value.(string), nil
	case int:
		return fmt.Sprintf("%d", value), nil
	case float64:
		return fmt.Sprintf("%f", value), nil
	case float32:
		return fmt.Sprintf("%f", value), nil
	default:
		return "", fmt.Errorf("%v无法转换成字符串", reflect.TypeOf(value))
	}
}

// InterfaceIsEmpty 是否是空
func InterfaceIsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	switch value.(type) {
	case string:
		if value.(string) != "" {
			return false
		}
	case int:
		if value.(int) != 0 {
			return false
		}
	case float64:
		if value.(float64) != 0.0 {
			return false
		}
	case float32:
		if value.(float64) != 0.0 {
			return false
		}
	default:
	}
	return true
}
