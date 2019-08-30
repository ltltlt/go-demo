package main

import "net/http"

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
	if r.FormValue("a") == "1" {
		panic(1)
	}
}
func main() {
	http.HandleFunc("/test", test)
	http.ListenAndServe(":4000", http.DefaultServeMux)
}
