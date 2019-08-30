package main

import "fmt"

func main() {
	m := make(map[int]string)
	c := make(chan struct{})

	go func() {
		for i := 0; i < 10000; i++ {
			m[i] = "abc"
		}
		c <- struct{}{}
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			fmt.Println(m[i])
		}
		c <- struct{}{}
	}()
	<-c
	<-c
}
