package util_test

import (
	"log"
	"testing"

	"github.com/mumushuiding/util"
)

func TestTimeStrSub(t *testing.T) {
	s := "abABcUser"
	f, _ := util.IsABC(s)

	log.Println(f)
}
