package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var s string = "true"
	//将 string true 转换为 boolean true
	b, _ := strconv.ParseBool(s)
	fmt.Println(b)

	//将 boolean true 转换为 string true
	boolString := strconv.FormatBool(b)
	fmt.Println("value ", boolString)
	fmt.Println("type ", reflect.TypeOf(boolString))

}
