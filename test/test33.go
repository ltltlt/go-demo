package main

func main() {
	foo(nil)
}

func foo(c []byte) {
	_ = c[2]
}
