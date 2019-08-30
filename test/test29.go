package main

type S struct{}

func main() {
	_ = identity()
}

func identity() *S {
	return &S{}
}
