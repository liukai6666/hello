package main

import (
	"fmt"
)

type DemoStruct struct {
	name string
	no   int
}

func main() {
	s := "HELLO WORLD!"
	t := "GoodBye,Cruel WORLD"
	fmt.Println("string is %v,t is %v\n", s, t)

	/**
	 * 	结合使用动词v和+ 来打印结构体中字段的名称。
	 */
	demoStruct := DemoStruct{
		name: "刘凯",
		no:   1,
	}
	fmt.Printf("%+v\n", demoStruct)
	fmt.Printf("%v\n", demoStruct)

}
