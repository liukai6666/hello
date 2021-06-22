/*
	if、else、else if
	使用比较运算符
	使用逻辑运算符
	使用switch语句
	使用for语句执行循环
	使用defer语句
*/
package main

import (
	"fmt"
	// "reflect"
	"strconv"
)

func main() {

	//if
	boolVariable := true
	if boolVariable {
		outputStr("b is true")
	}

	//if else
	boolVariableFalse := false
	if boolVariableFalse {
		outputStr("b is true")
	} else {
		outputStr("b is false")
	}

	//if elseif
	intVariable := 2
	if intVariable == 3 {
		outputStr("i equal 3")
	} else if intVariable == 2 {
		outputStr("i equal 2")
	}

	/**
	 *
		使用比较运算符
		两个字符串是否相同？ ==
		两个数字是否相同？ ==
		一个数字是否大于另一个数字？ >
		一个数字是否小于等于另一个数字？ <=
		关于Go语言的比较运算符，两个操作数的类型必须相同
		与 && 或 || 非 ！
	*/

	/**
	 *	switch
	 */

	intVariable2 := 4
	switch intVariable2 {
	case 2:
		outputStr("Two")
	case 3:
		outputStr("Three")
	case 4:
		outputStr("Four")
	}

	/**
	 *	switch default
	 */
	stringVariable := "c"
	switch stringVariable {
	case "a":
		outputStr("This's a")
	case "b":
		outputStr("This's b")
	default:
		outputStr("i don't recongize that letter!")
	}

	fmt.Println()
	//for
	i := 0
	for i < 10 {
		i++
		fmt.Println("i =", i)
	}

	fmt.Println()
	for j := 0; j < 5; j++ {
		fmt.Println("j =", j)
	}

	fmt.Println()
	for k := 0; k < 10; {

		fmt.Println("k =", k)
		k++
	}
	fmt.Println()

	//range 遍历数组、 这个像迭代器啊
	numbers := []string{"one", "two", "three", "four"}
	for index, num := range numbers {

		outputStr("index is " + strconv.FormatInt(int64(index), 10) + "，value equal \"" + num + "\"")
	}

}

func outputStr(str string) {
	fmt.Println(str)
	fmt.Println()
}
