/**
 * 	自定义结构体数据字段的默认值
 * 	    创建数据结构时，自定义数据字段的默认值是很有必要的。默认情况下，Go给数据字段指定相应数据类型的零值。
 *      默认零值
 *      Boolean false
 *      Integer 0
 *      Float 0.0
 *      String ""
 *      Pointer nil
 *      Function    nil
 *      Interface   nil
 *      Slice   nil
 *      Channel nil
 *      Map     nil
 *      Go语言没有提供自定义默认值的内置方法，但可使用构造函数来实现这个目标。构造函数创建结构体，并将没有指定值的数据字段设置为默认值。
 */
package main

import (
	"fmt"
)

type Alarm struct {
	Time  string
	Sound string
}

func main() {

	na := NewAlarm("2021-06-23 20:38:00")
	fmt.Printf("%+v\n", na)
}

/**
 *  也太lowB了吧
 */
func NewAlarm(time string) (a Alarm) {
	a = Alarm{
		Time:  time,
		Sound: "Klaxon",
	}

	return
}
