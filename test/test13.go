package main

import "fmt"

func main() {
	a := 11
	{
		a := a         // weird but legal(illegal in python, unknown result in C/C++)
		fmt.Println(a) // get 11
	}
}
