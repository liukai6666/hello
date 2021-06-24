/**
 *	创建字符串字面量
 * 	理解rune字面量
 * 	拼接字符串
 * 	编码
 *
 * 	特殊的字符
 * 	\a 	alert or bell
 * 	\b 	退格
 * 	\f 	换页符
 * 	\n 	换行符
 * 	\r 	回车
 * 	\t 	水平制表符
 * 	\v	垂直制表符
 *	\\	反斜杠
 * 	\'	单引号
 * 	\\" 双引号
 *
 * 	理解字符串是什么？
 * 		要理解字符串是什么，必须明白计算机是如何显示和存储字符的。计算机将数据解读为数字，另外，虽然您根本看不到，但计算机实际上是将字符存储为数字的。
 * 	历史上有很多编码标准，最后大家就如何将字符映射到数字达成了一致。ASCII曾经是最重要的编码标准，它就如何使用数字来表示英语字母表中字符进行了标准化。
 *		ASCII编码标准定义了如何使用7位的整数（通俗地说就是数字）来表示128个字符。虽然ASCII在英语字符标准化的道路上迈出了重要的一步，但它不包含其他语言的字符集。
 * 简而言之，它支持使用英语说"hello"，但不支持使用中文说"您好"
 * 		鉴于此，Unicode编码方案于1987年应运而生，它支持全球各地的大多数字符集。
 * 		很多的字符编码方案都实现了Unicode标准，其中最著名的就是UTF-8
 * 		要更深入地理解字符串以及如何操作它们，必须首先知道Go语言的字符串实际上是只读的字节切片。要获悉字符串包含多少个字节，可使用Go语言的内置函数len
 *
 */
package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "I am an interpreted string literal"
	fmt.Println(s)

	/**
	 * 	原始字符串，用反引号，特殊字符在反引号中不起任何作用。
	 */
	s = `    After a backslash, certain single character escapes represent 
special values n is a line feed or new line 
    t is a tab`
	fmt.Println(s)

	/**
	 *	拼接字符串
	 * 	在Go语言中，要拼接（合并）字符串，可将运算符+用于字符串变量。
	 */
	s = "hello, " + "liukai"
	fmt.Println(s)

	/**
	 * 使用复合赋值运算符 += 来拼接字符串
	 * 只能拼接类型为字符串的变量
	 */
	s = "Can you hear me?"
	s += "\nHear me screamin?"
	fmt.Println(s)

	/**
	 * 	若将整型和字符串拼接，会产生报错，需要先将整型转换为字符串后，再进行拼接操作
	 */
	num := 1
	str := "已经不是整型1了，此时1是\"1\""
	//	numStr := strconv.FormatInt(int64(num), 10)
	numStr := strconv.Itoa(num)
	fmt.Println(numStr + str)

	/**
	 *	使用缓冲区拼接字符串
	 * 	对于简单而少量的拼接，使用运算符+和+=的效果虽然很好，但随着拼接操作次数的增加，这种做法的效率并不高。
	 * 	如果在循环中拼接字符串，则使用空的字节缓冲区来拼接的效率更高
	 *
	 * 	循环500次，生成字符串
	 */
	var buffer bytes.Buffer
	for i := 0; i < 500; i++ {
		buffer.WriteString("z")
	}
	fmt.Println(buffer.String())

	s = "hello"
	//字符显示 h
	fmt.Printf("%q\n", s[0])
	//二进制显示
	fmt.Printf("%b\n", s[0])
	//显示的是10进制的104
	fmt.Println(s[0])

	/**
	 *	处理字符串
	 * 	给字符串变量赋值后，就可以使用标准库中的strings包提供的任何方法。这个包提供了一套完备的字符串处理函数，其文档非常的详尽。
	 */
	/**
	 *	将字符串转换为小写
	 * 	ToLower
	 */
	upperString := "ABCD"
	fmt.Println("ABCD TO LOWER " + strings.ToLower(upperString))

	/**
	 * 在字符串中查找子串
	 * Index
	 * 第一个参数：源字符串
	 * 第二个参数：要查找的子串，如果找到就返回第一个子串出现的索引位置，如果没有找到就返回-1
	 */
	fmt.Println("CD在ABCD中出现的位置是 ：" + strconv.Itoa(strings.Index(upperString, "CD")))

	/**
	 * 	删除字符串中的开头和末尾的空格
	 */
	trimSpaceStr := strings.TrimSpace("        abc             ")
	fmt.Println(trimSpaceStr)
	fmt.Println(len(trimSpaceStr))
}
