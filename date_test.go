package util_test

import (
	"fmt"
	"testing"

	"github.com/mumushuiding/util"
)

func TestTimeStrSub(t *testing.T) {
	str1 := "2019-08-21 12:56:01"
	str2 := "2019-08-21 12:57:02"
	result, err := util.TimeStrSub(str1, str2, util.YYYY_MM_DD_HH_MM_SS)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
