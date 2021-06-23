/**
 *	数组是一个数据集合，常用于存储一系列用数字做索引的数据
 * 	数组声明、声明一个数组变量，并指定其长度和数据类型
 * 		var cheers [2]string 数组声明
 * 		cheers[0] = "Mariolles"	数组赋值
 * 		cheers[1] = "Epoisses de Bour gogne"
 * 		fmt.Println(cheers[0])	打印数组某个元素
 * 		fmt.Println(cheers)	打印数组所有元素
 */
package main

import (
	"fmt"
)

func main() {
	var names [2]string
	names[0] = "liukai"
	names[1] = "liukai1"
	//./array.go:20:7: invalid array index 2 (out of bounds for 2-element array)
	//	names[2] = "liukai2"
	//[liukai liukai1]
	fmt.Println(names)
	fmt.Println(names[1])
}
