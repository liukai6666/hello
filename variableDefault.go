package main

import (
	"fmt"
)

var sss = "123"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q \n", i, f, b, s)

	if s == "" {
		fmt.Println("s is ''")
	}

	s1 := "Hello World"
	fmt.Println(s1)
	fmt.Println(sss)
}
