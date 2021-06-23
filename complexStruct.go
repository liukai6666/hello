/**
 * 	嵌套结构体
 * 		有时候，数据结构需要包含多个层级。此时，虽然可选择使用诸如切片等数据类型，但有时候需要的数据结构更复杂。为建立较为复杂的数据结构，在一个结构体中嵌套另一个结构体的方式很有用。
 */

package main

import (
	"fmt"
)

func main() {

	type Address struct {
		Number int
		Street string
		City   string
	}

	type Superhero struct {
		Name    string
		Age     int
		Address Address
	}

	m := Superhero{
		Name: "liukai",
		Age:  18,
		Address: Address{
			Number: 1,
			Street: "朝阳望京南湖中园",
			City:   "北京",
		},
	}

	//{Name:liukai Age:18 Address:{Number:1 Street:朝阳望京南湖中园 City:北京}}
	fmt.Printf("%+v\n", m)
	fmt.Println(m.Address.Street)
}
