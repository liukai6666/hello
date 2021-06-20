/*
指针：变量在内存中的地址，可在变量名前加上&字符。
*/
package main

import (
	"fmt"
)

func main() {

	s := "a"
	fmt.Println("变量在内存中的地址是： ", &s)
}
