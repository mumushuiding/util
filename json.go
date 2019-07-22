package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ToJSONStr 对象转换成字符串
func ToJSONStr(data interface{}) (string, error) {
	result, err := json.Marshal(data)
	return fmt.Sprintf("%s", result), err
}

// ToPageJSON 转换成json字符串
func ToPageJSON(datas interface{}, count, pageIndex, pageSize int) (string, error) {
	data, err := json.Marshal(datas)
	result := fmt.Sprintf("{\"rows\":%s,\"pageSize\":%d,\"total\":%d,\"page\":%d}", data, pageSize, count, pageIndex)
	return result, err
}

// Str2Struct Str2Struct
// 字符串转对象
func Str2Struct(source string, destination interface{}) error {
	err := json.Unmarshal([]byte(source), destination)
	return err
}

// Body2Map 获取前台传递的body值，并转化成map
func Body2Map(request *http.Request) (map[string]interface{}, error) {
	s, _ := ioutil.ReadAll(request.Body)
	if len(s) == 0 {
		return nil, nil
	}
	map1 := make(map[string]interface{})
	err := json.Unmarshal(s, &map1)
	if err != nil {
		return nil, err
	}
	return map1, nil
}

// Body2Struct 获取前台传递的body值，并转化成指定结构体
func Body2Struct(r *http.Request, pojo interface{}) (err error) {
	s, _ := ioutil.ReadAll(r.Body)
	if len(s) == 0 {
		return nil
	}
	err = json.Unmarshal(s, pojo)
	return err
}

// GetBody 获取 body 值，并存储到指定 struct 中去
func GetBody(result interface{}, w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return errors.New("只支持POST方法")
	}
	if err := Body2Struct(r, result); err != nil {
		return err
	}
	return nil
}
