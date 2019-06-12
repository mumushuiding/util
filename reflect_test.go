package util_test

import (
	"fmt"
	"testing"

	"github.com/mumushuiding/util"
)

type Configuration struct {
	Port           string `json:"SERVER_PORT"`
	ReadTimeout    int
	WriteTimeout   int
	DbLogMode      bool   `json:"DB_LOG_MODE"`
	DbType         string `json:"DB_TYPE"`
	DbName         string `json:"DB_NAME"`
	DbHost         string `json:"DB_HOST"`
	DbPort         string `json:"DB_PORT"`
	DbUser         string `json:"DB_USER"`
	DbPassword     string `json:"DB_PASSWORD"`
	DbMaxIdleConns int    `json:"DB_MaxIdleConns"`
	DbMaxOpenConns int    `json:"DB_MaxOpenConns"`
}

// func TestGetTagJSONArrayFromStruct(t *testing.T) {
// 	// Configuration 数据库配置结构
// 	result, err := util.GetTagJSONArrayFromStruct(&Configuration{})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)
// }
func BenchmarkGetTagJSONArrayFromStruct(b *testing.B) {
	// Configuration 数据库配置结构

	for i := 0; i < b.N; i++ {
		_, err := util.GetTagJSONArrayFromStruct(&Configuration{})
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(result)
	}
}

func TestGetTagJSONChannelFromStruct(t *testing.T) {
	resultStream, err := util.GetTagJSONChannelFromStruct(&Configuration{})
	if err != nil {
		fmt.Println(err)
	}
	for v := range resultStream {
		fmt.Println(v)
	}
	fmt.Println("结束")
}
func BenchmarkGetTagJSONChannelFromStruct(b *testing.B) {

	for i := 0; i < b.N; i++ {
		resultStream, err := util.GetTagJSONChannelFromStruct(&Configuration{})
		if err != nil {
			fmt.Println(err)
		}
		for _ = range resultStream {
			// fmt.Println(v)
		}
		// fmt.Println("结束")
	}
}
func TestGetFieldChannelFromStruct(t *testing.T) {
	resultStream, _ := util.GetFieldChannelFromStruct(&Configuration{})
	for v := range resultStream {
		fmt.Println(v)
	}
}
func TestStructSetValByReflect(t *testing.T) {
	var config = &Configuration{}
	err := util.StructSetValByReflect(config, "DB_NAME", 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	str, _ := util.ToJSONStr(config)
	fmt.Println(str)
}
