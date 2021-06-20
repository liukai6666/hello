/*
指针：变量在内存中的地址，可在变量名前加上&字符。
*/
package main

import (
	"fmt"
)

func main() {

	/*
		s := "a"
		//变量在内存中的地址是：  0xc00008e1e0
		fmt.Println("变量在内存中的地址是： ", &s)
	*/

	i := 1
	//i address is : 0xc0000b2008
	fmt.Println("i address is :", &i)
	//传值、内存新开辟空间
	showMemoryAddress(i)
}

func showMemoryAddress(x int) {
	//x address is : 0xc0000b2020
	fmt.Println("x address is :", &x)
	return
}
