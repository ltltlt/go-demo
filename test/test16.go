package main

//import "io/ioutil"

import "unicode/utf8"
import "fmt"
import "os"

func main() {
	f, _ := os.Open("/home/ty-l1/桌面/source导入导出模板及相关字段校验说明.csv")
	var bs []byte = make([]byte, 1000000)
	f.Read(bs)
	for i, count := 0, 0; i < len(bs); i += count {
		_, count := utf8.DecodeRune(bs[i:])
		if !utf8.FullRune(bs[i : i+count]) {
			fmt.Println(bs[i : i+count])
		}
	}
}
