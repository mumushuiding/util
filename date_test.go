package util_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mumushuiding/Activiti-go/util"
)

func TestFormateDate(t *testing.T) {
	fmt.Println(util.FormatDate(time.Now(), util.YYYY_MM_DD_HH_MM_SS))
}
