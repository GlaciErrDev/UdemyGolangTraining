package main

import "fmt"

func foo() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := foo()
	fmt.Println(nextInt())
	//fmt.Println(nextInt())
}
