/**
 * 	结构体是什么？
 * 		结构体是一系列具有指定数据类型的数据字段，它能够让您通过单个变量引用一系列相关的值。通过使用结构体，可在单个变量中存储众多类型不同的数据字段。存储在结构体中的值可轻松地访问和
 * 	修改，这提供了一种灵活的数据结构创建方式。通过使用结构体，可提高模块化程度，还能够让您创建并传递复杂的数据结构。
 * 		可将结构体视为用于创建数据记录（如员工记录和机票预定）的模版。
 * 		在c和c++中，结构体很常见，作为c语言家族的一员，Go也支持结构体。结构体并非创建面向对象代码的方式，而是一种数据结构的创建方式，旨在满足数据建模需求
 *
 * 	创建结构体
 * 		1. var variable struct
 * 		2. 使用new关键字
 * 		3. 使用简短变量赋值来创建结构体实例（最常用及最推荐的方式）
 * 	嵌套结构体
 * 		complexStruct.go
 * 	自定义结构体数据字段的默认值
 * 		customizeStructDefaultVal.go
 * 	比较结构体
 * 		compareStruct.go
 *  理解公有和私有值
 *      如果一个值被导出，可在函数、方法或包外面使用，那么它就是公有的；如果一个值只能在其所属上下文中使用，那么它就是私有的。
 *      结构体及其数据字段都可能被导出，如果一个标识符的首字母是大写的，那么它将被导出；否则不会导出。要导出结构体及其字段，结构体及其字段的名称都必须以大写字母开头
 * 	区分指针引用和值引用
 *		使用结构体时，明确指针引用和值引用的区别很重要。
 * 		数据值存储在计算机内存中。指针包含值的内存地址，这意味着使用指针可读写存储的值。创建结构体实例时，给数据字段分配内存并给它们指定默认值；然后返回指向内存的指针，并
 * 		将其赋值给一个变量。使用简短变量赋值时，将分配内存并指定默认值	a := Drink{}
 * 		复制结构体时，明确内存方面的差别很重要。将指向结构体的变量赋值给另一个变量时，被称为赋值		a := b
 * 		赋值后，a与b相同，但它是b的副本，而不是指向b的引用。修改b不会影响a
 */
package main

import (
	"fmt"
)

/**
 * 	结构体的声明
 */
type Movie struct {
	Name   string
	Rating float32
}

func main() {

	var m Movie
	/**
	 * 	%+v = struct
	 *	创建结构体时，如果没有初始化，则Go将把每个数据字段设置为相应数据类型的零值。
	 * 	{Name: Rating:0}
	 */
	fmt.Printf("%+v\n", m)
	/**
	 *	结构体字段赋值
	 */
	m.Name = "liukai"
	m.Rating = 10.1
	fmt.Printf("%+v\n", m)

	m1 := Movie{
		Name:   "Citizen Kane",
		Rating: 10,
	}
	//{Name:Citizen Kane Rating:10}
	fmt.Printf("%+v\n", m1)

	/**
	 *	使用new关键字来创建结构体实例
	 */
	m2 := new(Movie)
	m2.Name = "liukai2"
	m2.Rating = 100.2
	//&{Name:liukai2 Rating:100.2}
	fmt.Printf("%+v\n", m2)

	/**
	 *	使用简短变量赋值来创建结构体实例
	 * 	这种方式为常用用的结构体创建方式、也是最推荐的方式
	 */
	m3 := Movie{Name: "liukai3", Rating: 9.99}
	fmt.Printf("%+v\n", m3)

	/**
	 * 	也可省略字段名，按字段声明顺序给它们赋值，但出于维护性考虑，不推荐这样做
	 */
	m4 := Movie{"liukai4", 6.66}
	fmt.Printf("%+v\n", m4)

	/**
	 *	字段很多时，让每一个字段独占一行能够提高代码的可维护性和可读性
	 */
	var m5 Movie = Movie{
		Name:   "liukai5",
		Rating: 6.666,
	}
	fmt.Printf("%+v\n", m5)

	am := Movie{
		Name:   "英雄本色",
		Rating: 10.00,
	}
	cpam := am
	cpam.Name = "英雄本色2"

	//am is {Name:英雄本色 Rating:10}
	fmt.Printf("am is %+v\n", am)
	//cpam is {Name:英雄本色2 Rating:10}
	fmt.Printf("cpam is %+v\n", cpam)
	//am address is 0xc00000c160
	fmt.Printf("am address is %p\n", &am)
	//cpam address is 0xc00000c180
	fmt.Printf("cpam address is %p\n", &cpam)

	/**
	 * 	要修改原始结构体实例包含的值，必须使用指针。指针是指向内存地址的引用，因此使用它操作的是结构体本身。要获得指针，可在变量名前加上&
	 */
	cpAmAddress := &am
	fmt.Printf("cpAmAddress is %p\n", cpAmAddress)
	//	改变该址指向的值
	//	./structural.go:118:2: invalid indirect of cpAmAddress.Name (type string)
	//	*cpAmAddress.Name = "英雄本色3"
	//	fmt.Printf("am is %+v", am)
	cpAmAddress.Name = "英雄本色3"
	fmt.Printf("am is %+v\n", am)

}
