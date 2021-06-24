/**
 * 	使用方法
 *      相对于面向对象来说 struct定义类的属性、方法定义类的方法
 *      方法类似于函数，但有一点不同：在关键字func后面添加了另一个参数部分，用于接受单个参数。
 *      type Movie struct {
 *          Name    string
 *          Rating  float32
 *      }
 *      func (m *Movie) summay() string {
 *          //code
 *      }
 *      在方法声明中，关键字func后面多了一个参数--接收者。严格地说，方法接收者是一种类型，这里是指向结构体Movie的指针。接下来是方法名、参数以及返回类型。除多了包含接收者的参数部分
 *      外，方法与函数完全相同。可将接收者视为与方法相关联的东西。通过声明方法Summary，让结构体Movie的任何实例都可使用它。
 *      为何要使用方法，而不直接使用函数呢？
 *      例如，下面的函数与前面的方法等价
 *      func summary(m *Movie) string {
 *          //code
 *      }
 *      函数summary和结构体Movie相互依赖，但它们之间没有直接关系。例如，如果不能访问结构体Movie的定义，就无法声明函数summary。如果使用函数，则在每个使用函数或结构体的地方，都
 *      需包含函数和结构体的定义，这会导致代码重复。另外，函数发生任何改变，都必须随之修改多个地方。这样看来在函数与结构体发生关系密切时，使用方法更合理。
 * 	创建方法集
 *      方法集是可对特定数据类型进行调用的一组方法。在Go语言中，任何数据类型都可有相关联的方法集，这让您能够在数据类型和方法之间建立关系
 *      方法集可包含的方法数量不受限制，这是一种封装功能和创建代码的有效方式
 * 	使用方法和指针
 *      与结构体一样，明白如何使用方法和指针也很重要。您已经看到，方法是一个接受被称为接收者的特殊参数的函数，接收者可以是指针，也可以是值，但两者的差别非常微妙
 *      指针和值之间的差别很微妙，但选择使用指针还是值这一点很简单，如果需要修改原始结构体，就使用指针；如果需要操作结构体，但不想修改原始结构体，就使用值
 */
package main

import (
	"fmt"
	"math"
	"strconv"
)

type Movie struct {
	Name   string
	Rating float64
}

/**
 *  创建方法
 *  在方法声明中，关键字func后面多了一个参数--接收者。严格地说，方法接收者是一种类型，这里是指向结构体Movie的指针。接下来是方法名、参数以及返回类型。除多了包含接收者的参数部分
 *  外，方法与函数完全相同。可将接收者视为与方法相关联的东西。通过声明方法Summary，让结构体Movie的任何实例都可使用它
 */
func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + "," + r
}

/**
 *  使用函数当然可以实现同样的功能，但是一个类会有多个方法，方法多了，函数写的多，此时就体现了封装的意义，故方法优于函数
 */
func summary(m *Movie) string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + "," + r

}

/**
 *  处理球体时，假设您要计算其表面积和体积。在这种情况下，非常适合使用结构体和方法集。通过使用方法集，您只需创建一次计算代码，就可将其重用于任何球体
 * 定义球体属性 -- 半径
 */
type Sphere struct {
	Radius float64
}

/**
 *  计算球体表面积
 */
func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}

func (s *Sphere) Volume() float64 {
	radiusCubed := s.Radius * s.Radius * s.Radius
	return (float64(4) / float64(3)) * math.Pi * radiusCubed
}

type Triangle struct {
	base   float64
	height float64
}

func (t *Triangle) area() float64 {
	return 0.5 * (t.base * t.height)
}

/**
 *  接收者为指针
 */
func (t *Triangle) changeBasePointer(f float64) {
	t.base = f
}

/**
 *  接收者非指针
 */
func (t Triangle) changeBaseNonPointer(f float64) {
	t.base = f
}

func main() {
	m := Movie{
		Name:   "义胆红唇",
		Rating: 9.11,
	}

	fmt.Println(m.summary())

	s := Sphere{
		Radius: 5,
	}
	fmt.Printf("s.SurfaceArea equal %v\n", s.SurfaceArea())
	fmt.Printf("s.Volume equal %v\n", s.Volume())

	t := Triangle{
		base:   3,
		height: 1,
	}
	fmt.Println(t.area())
	t.changeBasePointer(6)
	fmt.Printf("t.base changed 6, %+v\n", t)
	t.changeBaseNonPointer(3)
	fmt.Printf("t.base no change, %+v\n", t)

}
