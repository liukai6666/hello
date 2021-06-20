/*
	在go语言中，一对大括号{}表示一个块。
	对于在大括号{}内声明的变量，可在相应块的任何地方访问
	大括号内的大括号定义了一个新块---内部块
	在内部块中，可访问外部块中声明的变量
	在外部块中，不能访问在内部块中声明的变量
	简而言之，每个内部块都可访问其外部块，但外部块不能访问内部块
*/
package main

import (
	"fmt"
)

var s = "Hello World"

func main() {
	fmt.Printf("Print 's' variable from outer block %v\n", s)

	b := true
	if b {
		fmt.Printf("Printing 'b' varibale from outer block %v\n", b)

		i := 42
		if b != false {
			fmt.Printf("Printing 'i' varibale from outer block %v\n", i)
		}

		fmt.Printf("Printing 'i' varibale from inner block %v\n", i)
	}

	//	fmt.Printf("Printing 'i' varibale from inner block %v\n", i) //i 只能在内部块中访问到、此时报错
}
