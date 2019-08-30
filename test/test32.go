package main

import "fmt"
import "github.com/spf13/cast"

var m = make(map[string]string)
var ch = make(chan struct{})

func test() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover", r)
		}
	}()
	go read()
	go write()
	<-ch
	<-ch
}

func read() {
	for i := 0; i < 1000000; i++ {
		_ = m[cast.ToString(i)]
	}
	ch <- struct{}{}
}

func write() {
	for i := 0; i < 1000000; i++ {
		m[cast.ToString(i)] = cast.ToString(i + 1)
	}
	ch <- struct{}{}
}

func main() {
	test()
}
