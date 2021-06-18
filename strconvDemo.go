package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var s string = "true"
	b, _ := strconv.ParseBool(s)
	fmt.Println(b)

	boolString := strconv.FormatBool(b)
	fmt.Println("value ", boolString)
	fmt.Println("type ", reflect.TypeOf(boolString))

}
