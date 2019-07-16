package util_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestTest(t *testing.T) {
	fmt.Println("--------------------")
	groups := []string{"'日本'", "'中国'"}
	s := strings.Join(groups, ",")
	fmt.Println(s)
}
