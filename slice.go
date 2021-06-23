/**
 * 	切片是底层数组中的一个连续片段、通过它您可以访问该数组中一系列带编号的元素。因此，切片能够访问数组的特定的一部分
 * 	为何要使用切片？而不直接使用数组？
 * 		在Go语言中，使用数组存在一定的局限性，数组定长，您无法在数组中添加元素；而切片比数组灵活，您可以在切片中添加或者删除元素，还可以复制切片中的元素。切片既保留了数组的完整性
 * 		又比数组使用起来容易。
 * 	切片的声明
 * 		var cheers = make([]string, 2)  make 第一个餐数为类型、第二个参数为切片长度
 */
package main

import (
	"fmt"
)

func main() {

	//声明切片
	var names = make([]string, 2)
	names[0] = "liukai"
	names[1] = "liukai1"

	/**
	 * 	在切片中添加元素
	 * 	append函数只能用于切片
	 * 	用于数组报错： first argument to append must be slice; have [2]string
	 */
	names = append(names, "liukai2", "liukai3", "liukai4")
	fmt.Println(names)
	fmt.Println(len(names))
	/**
	 *	在切片中删除元素
	 * 	保留下标2之前的所有元素、不包含下标2
	 * 	append(names[:2])
	 */
	// names = append(names[:2])
	/**
	 * 	从下标2开始，保留之后的所有元素
	 * 	append(names[2:])
	 * 	索引将从新排序、由0开始
	 */
	names = append(names[2:])
	fmt.Println(names)
	fmt.Println(len(names))

	/**
	 *	将一个切片的元素复制到另一个切片中
	 * 	函数copy在新切片中创建元素的副本，因此修改一个切片中的元素不会影响另一个切片
	 */
	var names2 = make([]string, 3)
	copy(names2, names)

	fmt.Println("names2 : ", names2)
}
