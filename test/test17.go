package main

import "fmt"

func main() {
	s := "\u0001"
	b := []byte(s)
	fmt.Println(b)
}
