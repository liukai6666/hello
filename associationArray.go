/**
 *	数组和切片是通过索引值访问元素集合，而映射是通过键来访问无序元素编组。在其他编程语言中，映射也被称为关联数组、字典或散列。映射在信息查找方面的效率非常高，
 *  因为可直接通过键来检索数据。简单地说，映射可视为键值对集合。
 *  可在映射中动态地添加元素，而无须调整映射的长度
 */
package main

import (
	"fmt"
)

func main() {
	//关联数组的声明
	var myinformation = make(map[string]string)
	myinformation["name"] = "liukai"
	myinformation["age"] = "35"
	myinformation["sex"] = "male"
	myinformation["address"] = "北京市朝阳区望京南湖中园228楼5单元301"

	fmt.Println(myinformation["sex"])
	fmt.Println(myinformation["age"])

	/**
	 * 要从映射中删除元素，可使用内置函数delete
	 */
	delete(myinformation, "age")
	//map[address:北京市朝阳区望京南湖中园228楼5单元301 name:liukai sex:male]
	fmt.Println(myinformation)
}
