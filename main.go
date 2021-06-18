package main

import (
	"fmt"
)

func sayHello(name string) string {
	return "Hello " + name
}

func main() {
	//fmt.Println(sayHello("liukai"))
	fmt.Println(addition(1, 2))
}

func addition(x int, y int) int {
	return x + y
}
