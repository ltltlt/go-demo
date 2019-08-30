package main

//import "github.com/spf13/cast"
import "fmt"
import "encoding/json"

type Test struct {
	A int
	B int
}

func main() {
	t := Test{1, 2}
	bs, _ := json.Marshal(t)
	var m map[string]interface{}
	json.Unmarshal(bs, &m)
	fmt.Println(m)
}
