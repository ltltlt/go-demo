package main

func main() {
	test(0)
}

func test(i int) {
	//var array [1024 * 1024 * 1024]uint8
	//do_array(&array)
	i += 1
	//println(i)
	test(i)
}

func do_array(v *[1024 * 1024 * 1024]uint8) {
	for i := range v {
		v[i] = (uint8)(i*22 + 3)
	}
}
