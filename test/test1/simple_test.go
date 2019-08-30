package main

import "testing"
import "fmt"
import "github.com/kardianos/osext"

func Test(t *testing.T) {
	fmt.Println(osext.Executable())
}
