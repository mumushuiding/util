package util_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mumushuiding/util"
)

func TestFormateDate(t *testing.T) {
	fmt.Println(util.FormatDate(time.Now(), util.YYYY_MM_DD_HH_MM_SS))
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(s)
}
