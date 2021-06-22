/*
Go语言提供了一些函数式编程功能，如能够将一个函数作为参数传递给其他函数，Go语言将函数视为一种类型，因此可将函数赋值给变量
*/
package main

import (
	"fmt"
)

func main() {

	fn := func() {
		fmt.Println("function called")
	}

	fn()

	fn2 := func() string {
		return "function called"
	}

	fmt.Println(anotherFunction(fn2))
}

func anotherFunction(f func() string) string {
	return f() + "add anotherString"
}
