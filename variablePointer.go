/*
指针：变量在内存中的地址
使用指针变量名前加&  fmt.Println(&x)
声明指针，在变量名前加*  func transferPointer(x int*)
非参数声明指针部分、要访问指针地址指向的值，需要在变量前面加*号  x = pointer, fmt.Println(*x)
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
	//传址
	showMemoryAddressParameterPointer(&i)
}

func showMemoryAddress(x int) {
	//transfer value, x address is : 0xc0000b2020
	fmt.Println("transfer value, x address is :", &x)
	return
}

func showMemoryAddressParameterPointer(x *int) {

	//参数接收的就是地址、此时访问x就是访问地址、不需要再加&
	//transfer pointer, x address is : 0xc0000b2008
	fmt.Println("transfer pointer, x address is :", x)
	return
}
