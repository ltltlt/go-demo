package main

import "fmt"
import "time"

func main() {
	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(now.Location())
	fmt.Println(now)
	d := 24 * time.Hour
	fmt.Println(now.Truncate(d))
	fmt.Println(now.MarshalJSON())
}
