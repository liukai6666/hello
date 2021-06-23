package main

import (
	"fmt"
)

func main() {
	fmt.Println(`
该使用数组还是切片？
除非确定必须使用数组、否则请使用切片。切片能够轻松地添加和删除元素，还无须处理内存分配问题。

从切片中删除元素？
append（slice[2:]）、从下标2开始，保留之后的所有元素
append(slice[:2])、保留下标2之前的所有元素

需要指定映射的长度吗？
不需要。使用make创建映射时，可使用第二个参数，但这个参数只是容量提示，而非硬性规定。映射可根据要存储的元素数量自动增大，因此没有必要指定长度。
`)

	/**
	 *	创建一个包含4个元素的切片；再创建一个切片，并将第3个元素和第4个元素复制到该切片中
	 */
	var sli1 = make([]string, 4)
	sli1[0] = "a"
	sli1[1] = "b"
	sli1[2] = "c"
	sli1[3] = "d"
	fmt.Println(sli1)

	var sli2 = make([]string, 2)
	copy(sli2, append(sli1[2:]))

	//sli2 is  [c d]
	fmt.Println("sli2 is ", sli2)

	/**
	 * 	根据下述HTML元素列表创建一个映射
	 * 	p - 段落
	 * 	img - 图像
	 * 	h1 - 一级标题
	 * 	h2 - 二级标题
	 */
	var htmlInfo = make([string]string)
	htmlInfo["p"] = "段落"
	htmlInfo["img"] = "图像"
	htmlInfo["h1"] = "一级标题"
	htmlInfo["h2"] = "二级标题"
}
