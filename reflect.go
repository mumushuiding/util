package util

import (
	"errors"
	"reflect"
	"regexp"
	"strings"
)

// StructSetValByReflect StructSetValByReflect
// structname 是一个指针变量
func StructSetValByReflect(structname interface{}, fieldname string, value interface{}) error {
	pp := reflect.ValueOf(structname)         // 取得struct变量的指针
	field := pp.Elem().FieldByName(fieldname) // 获取指定field
	if !field.IsValid() {
		return errors.New("结构体没有属性【" + fieldname + "】")
	}
	// fmt.Printf("------------%v", field)
	valueType := reflect.TypeOf(value).String()
	if field.Type().String() != valueType {
		return errors.New("field的类型为【" + field.Type().String() + "】,value的类型为【" + valueType + "】,两者不匹配")
	}
	field.Set(reflect.ValueOf(value))
	// switch valueType {
	// case "string":
	// 	field.SetString(value.(string))
	// 	break
	// case "int":
	// 	fi
	// 	field.SetInt(value.(int))
	// 	break
	// case "bool":
	// 	field.SetBool(value.(bool))
	// 	break
	// }
	return nil
}

// GetTagJSONChannelFromStruct GetTagJSONChannelFromStruct
// 当属性比较多，后序操作比较复杂的时候，建议使用
func GetTagJSONChannelFromStruct(structName interface{}) (<-chan string, error) {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, errors.New("不是结构体，not Struct")
	}
	fieldNum := t.NumField()
	if fieldNum == 0 {
		return nil, errors.New("结构体没有任何属性")
	}
	resultStream := make(chan string, fieldNum)
	reg := regexp.MustCompile(`.*json:"([^"]*)".*`)
	go func() {
		defer close(resultStream)
		// defer fmt.Println("关闭resultStream")
		for i := 0; i < fieldNum; i++ {
			var result string
			tags := string(t.Field(i).Tag)
			if strings.Contains(tags, "json:") {
				result = reg.ReplaceAllString(tags, "$1")
			} else {
				result = t.Field(i).Name
			}
			resultStream <- result
			// fmt.Printf("send %s\n", result)
		}
	}()
	return resultStream, nil
}

// GetTagJSONArrayFromStruct GetTagJSONArrayFromStruct
// 当属性比较少，后序没有什么操作时使用
func GetTagJSONArrayFromStruct(structName interface{}) ([]string, error) {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, errors.New("不是结构体，not Struct")
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	reg := regexp.MustCompile(`.*json:"([^"]*)".*`)
	for i := 0; i < fieldNum; i++ {
		tags := string(t.Field(i).Tag)
		if strings.Contains(tags, "json:") {
			result = append(result, reg.ReplaceAllString(tags, "$1"))
		} else {
			result = append(result, t.Field(i).Name)
		}
	}
	return result, nil
}

// GetFieldArrayFromStruct GetFieldArrayFromStruct
// 获取结构体的属性[]string
func GetFieldArrayFromStruct(structName interface{}) ([]string, error) {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, errors.New("不是结构体，not Struct")
	}
	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	// reg := regexp.MustCompile(`.*json:"([^"]*)".*`)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result, nil
}

// GetFieldChannelFromStruct GetFieldChannelFromStruct
// 当属性比较多，后序操作比较复杂的时候，建议使用
func GetFieldChannelFromStruct(structName interface{}) (<-chan string, error) {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, errors.New("不是结构体，not Struct")
	}
	fieldNum := t.NumField()
	if fieldNum == 0 {
		return nil, errors.New("结构体没有任何属性")
	}
	resultStream := make(chan string, fieldNum)
	go func() {
		defer close(resultStream)
		// defer fmt.Println("关闭resultStream")
		for i := 0; i < fieldNum; i++ {
			resultStream <- t.Field(i).Name
			// fmt.Printf("send %s\n", result)
		}
	}()
	return resultStream, nil
}
