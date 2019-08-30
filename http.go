package main

import (
	"context"
	"fmt"
	"net/http"
)

func test(a int) int {
	return 1 + a
}

func main() {
	http.ListenAndServe(":9090", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%+v\n", r)
		ctx := context.Background()
		ctx = context.WithValue(ctx, "abc", test)
		fmt.Printf("%+v\n", r)
	}))
}
