/**
 * 	比较结构体
 *  对结构体进行比较，要先看它们的类型和值是否相同，对于类型相同的结构体，可使用相等性运算符来比较。要判断两个结构体是否相等，可使用==；要判断它们是否不等，可使用!=。
 */
package main

import (
	"fmt"
	"reflect"
)

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}

	b := Drink{
		Name: "Lemonade",
		Ice:  true,
	}

	if a == b {
		fmt.Println("a and b the same")
	}

	fmt.Printf("%+v\n", a)
	//main.Drink
	fmt.Println(reflect.TypeOf(a))
	fmt.Printf("%+v\n", b)
	//main.Drink
	fmt.Println(reflect.TypeOf(b))
}
