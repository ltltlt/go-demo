package main

import "os"
import "fmt"

func main() {
	fmt.Println(os.Readlink("/proc/self/exe"))
}
