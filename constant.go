/*
常量在程序生命周期中，只可访问值，不可修改值
*/
package main

import (
	"fmt"
)

const greeting string = "Hello,World"

func main() {
	//./constant.go:13:11: cannot assign to greeting、常量不能重新赋值
	//	greeting = "Bye,Curel world"
	fmt.Println(greeting)
}
