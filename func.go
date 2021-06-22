package main

import (
	"fmt"
	"strconv"
)

func main() {
	n1 := 1
	n2 := 2
	fmt.Println("Result : ", addUp(n1, n2))
	fmt.Println()

	rs := isOdd(n1)
	if rs == true {
		fmt.Println(n1, "是奇数")
	} else {
		fmt.Println(n1, "是偶数")
	}
	fmt.Println()

	_, goodName := getPrize()
	fmt.Println("商品名称：", goodName)
	fmt.Println()

	result := sumNumbers(1, 2, 3, 4)
	fmt.Println("1,2,3,4 Sum Result : ", result)
	fmt.Println()

	str1, str2 := sayHi()
	fmt.Println(str1, str2)
	fmt.Println("")

	fmt.Println("5的阶乘结果 ：", factorial(5, 0))
	fmt.Println()

	feedMe(1, 1)
	fmt.Println()
}

func addUp(x int, y int) int {
	return x + y
}

//单个输入、单个输出、 判断一个数字是奇数还是偶数
func isOdd(number int) bool {
	return number%2 == 1
}

//多个输出
func getPrize() (int, string) {
	return 2, "金鱼"
}

//定义不定参数函数
/*
不定参数函数是参数数量不确定的函数。通俗地说，这意味着它们接受可变数量的参数。在Go语言中，能够传递可变数量的参数，但它们的类型必须与函数签名指定的类型相同。要指定不定参数，可使用
三个点...。
*/
func sumNumbers(numbers ...int) int {

	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

/*
	使用具名返回值
	这个函数体中，在终止语句return前给具名变量进行了赋值。使用具名返回值时，无须显示地返回相应的变量。这被称为naked return语句
*/
func sayHi() (x, y string) {
	x = "Hello"
	y = "World"
	return
}

/*
	递归函数 阶乘
*/
var factorialString string = ""

func factorial(num int, frequency int) (result int) {

	if num > 1 {

		frequency = frequency + 1
		result = num * factorial(num-1, frequency)

		factorialString = factorialString + " * " + strconv.FormatInt(int64(num), 10)
		if frequency == 1 {
			factorialString = factorialString + " = " + strconv.FormatInt(int64(result), 10)
			fmt.Println(factorialString)
		}
		return
	}

	factorialString = factorialString + strconv.FormatInt(int64(1), 10)

	result = 1
	return
}

/**
 *	递归函数2
 */
func feedMe(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten > 5 {
		fmt.Printf("i'm full! i've eaten %d\n", eaten)
		return eaten
	}
	fmt.Printf("i 'm still hungry! i 've eaten %d\n", eaten)
	return feedMe(portion, eaten)
}
