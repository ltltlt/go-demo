package more_test

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	var v [3]int
	var v1 [3]int
	v1[0] = 2
	v = v1
	fmt.Println(v)
}