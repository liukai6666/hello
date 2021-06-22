package main

import (
	"fmt"
)

func main() {

	/**
	 *	使用defer语句
	 *	在函数返回前执行另一个函数
	 * 	函数在遇到return语句或到达函数末尾时返回。defer语句通常用于执行清理操作或确保操作（如网络调用）完成后再执行另一个函数
	 * 	类似面向对象的析构函数、做收尾工作
	 *	当存在多条defer时，先执行最后一条defer，先进后出
	 * 	执行顺序：
	 * 	i'm the third defer statement
	 *	i'm the second defer statement
	 *	I'm the first defer statement
	 *	I'm run after the function completes
	 */
	defer fmt.Println("I'm run after the function completes")

	defer fmt.Println("I'm the first defer statement")

	defer fmt.Println("i'm the second defer statement")

	defer fmt.Println("i'm the third defer statement")

	fmt.Println("Hello World")

}
