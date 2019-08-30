package main

import "time"
import "fmt"

func main() {
	fmt.Println(time.Now().Unix())
	time.Sleep(time.Second * 10)
	fmt.Println(time.Now().Unix())
}
