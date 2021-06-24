/**
 * 	在Go语言中，接口指定了一个方法集，这是实现模块化的强大方式。您可将接口视为方法集的蓝本，它描述了方法集中的所有方法，但没有实现它们。接口功能强大，因为它充当了方法集规范，这意味着
 * 	可在符合接口要求的前提下随意更换实现。
 * 	接口描述了方法集中的所有方法，并指定了每个方法的函数签名。
 *
 * 	假设需要编写一些控制机器人（Robot）的代码。粗略的说，可假定有多种类型的机器人，控制这些机器人行为的方式存在细微的差别。给定这个编程任务，您可能认为需要为每种机器人编写不同的
 * 	代码。通过使用接口，可将代码重用于有相同行为的实体。就这个机器人示例而言，下面的接口描述了开关机器人的方式。
 *
 * 	Go语言是面向对象吗？
 * 	Go语言没有提供类和类继承等面向对象的功能，但结构体和方法集弥补了这部分的不足，提供了一些面向对象元素。
 *
 * 	接口的意义
 * 	有助于代码重用，还能完全更换实现。假设编写一个使用MySQL数据库的计算机程序，如果不使用接口，则代码将完全是针对MySQL的。在这种情况下，如果后来要将MySQL数据库换成其他数据库，
 * 	如PostgreSQL，就可能需要重写大量的代码。通过定义一个数据库接口，该接口的实现将比使用的数据库更重要。从理论上说，只要实现满足接口的要求，就可以使用任何数据库，因此可轻松地更换
 * 	数据库。数据库接口包含多个实现，这就引入了多态的概念。
 * 	如果一个方法集实现了一个接口，就可以说它与另一个实现了该接口的方法集互为多态
 *
 * 	函数和方法的区别
 * 	方法多了一个指定接收者的参数
 */
package main

import (
	"errors"
	"fmt"
)

//机器人Robot接口、定义打开开关方法PowerOn，返回类型为error
type Robot interface {
	PowerOn() error
	Talk(str string)
}

//T850结构体、包含一个字服串型的Name属性
type T850 struct {
	Name string
}

//T850 的开机方法 必须实现接口Robot的PowerOn方法 、直接返回空值就行
func (t *T850) PowerOn() error {
	//永远机器人启动成功、也说明 error类型的默认值为nil
	return nil
}

func (t *T850) Talk(s string) {
	fmt.Println("hi,", t.Name, s)
}

//R2D2 结构体 、包含属性 Broken 值为bool类型
type R2D2 struct {
	Broken bool
}

//R2D2的开机方法，必须实现接口Robot的PowerOn方法， 如果R2D2坏了， 返回一个error R2D2 is broken、否则返回空
func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}

func (r *R2D2) Talk(s string) {
	fmt.Println("I'm R2D2", s)
}

//启动函数 参数为一个机器人、执行机器人的启动方法、返回一个错误类型
func Boot(r Robot) error {
	r.Talk("666")
	return r.PowerOn()
}

func main() {

	//实例话 T850、Name为The Terminaor
	t := T850{
		Name: "The Terminaor",
	}
	//实例话R2D2、Broken 设置为true
	r := R2D2{
		Broken: true,
	}

	//调用启动函数、将R2D2传入
	err := Boot(&r)
	//如果启动失败 打印错误信息、否则打印 开机成功
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is Powered on")
	}

	/**
	 * 	启动T850
	 * 	参数必须传址、否则报错cannot use t (type T850) as type Robot in argument to Boot:	T850 does not implement Robot (PowerOn method has pointer receiver)
	 */
	err = Boot(&t)
	//如果启动失败、打印错误信息、否则打印开机成功
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on")
	}
}
