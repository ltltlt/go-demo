package main

func test() {
	defer func() {
		return

		recover()
	}()

	panic("ooops")
}

func main() {
	test()
}
