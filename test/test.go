package main

import "bytes"
import "fmt"

func main() {
	bs := bytes.NewBuffer(make([]byte, 1024))
	bs.WriteString("你好")
	fmt.Println(bs.String())
}
