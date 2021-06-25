/**
 * 	错误处理及Go语言的独特之处
 * 	理解错误类型
 * 	创建错误
 * 	设置错误格式
 * 	从函数返回错误
 * 	错误和可用性
 * 	慎用panic
 * 	软件不可避免地会有错误及遇到未考虑到的情形，很多语言选择在未发生必须捕获的错误时引发异常，而Go语言处理错误的方式很有趣，将错误作为一种类型，这意味着可将错误
 * 	传递给函数和方法。
 *
 * 	Go语言中，一种约定是在调用可能出现问题的方法或函数时，返回一个类型为错误的值。这意味着如果出现问题，函数不会引发异常，而是让调用者决定如何处理错误
 *
 * 	理解错误类型
 * 	在Go语言中，错误是一个值。标准库声明了接口error
 * 	type error interface {
 * 		Error() string
 *  }
 * 	这个接口只有一个方法--Error，它返回一个字符串
 *
 *
 */
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
)

/**
 * 	从函数返回错误
 * 	Go语言的做法是从函数和方法中返回一个错误值。前面介绍了如何创建并返回错误
 */
func Half(numberToHalf int) (halfNumber int, err error) {
	if numberToHalf%2 != 0 {
		halfNumber = -1
		err = fmt.Errorf("Cannot half %v", numberToHalf)
	} else {
		halfNumber = numberToHalf / 2
		err = nil
	}

	return
}

func main() {

	/**
	 *	在Go语言中，有一种约定是，如果没有发生错误，返回的错误值将为nil。这让程序员调用方法或函数时，能够检查它是否像预期那样执行完毕。
	 *	在Go程序中，这种做法很常见。有些开发人员认为，这种做法很繁琐，因为它要求调用每个方法或函数时都检查错误，导致代码重复。
	 * 	这说得也许没错，但Go语言处理错误的方式比其他语言更灵活，因为可像其他类型一样在函数之间传递错误。这通常意味着代码要简短得多。
	 */
	file, err := ioutil.ReadFile("foo.txt")

	if err != nil {
		//*os.PathError
		fmt.Println(reflect.TypeOf(err))
		//open foo.txt: no such file or directory
		fmt.Println(err)
	}

	fmt.Println("%s", file)

	/**
	 *	您已经知道如何使用错误，但如果要创建并返回错误，该怎么办呢？标准库中的errors包支持创建和操作错误。
	 *	创建并打印错误的例子
	 */
	err = errors.New("Something went wrong")
	if err != nil {
		//*errors.errorString
		fmt.Println(reflect.TypeOf(err))
		//Something went wrong
		fmt.Println(err)
	}

	/**
	 *	设置错误的格式
	 * 	除errors包外，标准库中的fmt包还提供了Errorf，可用于设置返回的错误的字符串的格式，这能够让您将多个值合并称更有意义的错误字符串，从而动态地创建错误字符串。
	 * 	fmt包设置错误字符串的格式
	 */
	name, role := "RichardJupp", "Drummer"
	//*errors.errorString
	err = fmt.Errorf("The %v %v quit", role, name)
	fmt.Println(reflect.TypeOf(err))
	fmt.Println(err)

	n, err := Half(19)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	/**
	 *	错误和可用性
	 * 	除从技术角度考虑Go语言的错误处理方式和错误生成方式外，还需从以用户为中心的角度考虑错误。编写供他人使用的库或包时，您编写和使用错误的方式将极大地影响可用性。使用您编写的库
	 * 	时，用户可能遇到错误，并尝试从错误中恢复。
	 * 	如果库用户相信错误会以一致的方式返回，且包含有用的错误消息，则用户能够从错误中恢复的可能性将高得多。他们很可能会认为您编写的库不仅很实用，而且值得信息。
	 */
	/**
	 * 糟糕的错误信息
	 * You broke something! Good luck!
	 */
	err = errors.New("You broke something! Good luck!")
	fmt.Println(err)
	/**
	 * 	No config file found. Please create one at ~/.foorc
	 * 	这条错误信息比上面那一条更能让用户知道，具体是哪里出现了错误
	 * 	为什么这样说呢？有如下几个原因
	 * 		具体地指出了问题所在。
	 * 		提供了问题解决方案。
	 * 		对用户更有礼貌。
	 */
	err = errors.New("No config file found. Please create one at ~/.foorc")
	fmt.Println(err)

	/**
	 *	慎用panic
	 * 	panic是Go语言中的一个内置函数，它终止正常的控制流程并引发恐慌（panicking），导致程序停止执行。出现普通错误时，并不提倡这种做法，因为程序将停止执行，并且没有任何回旋
	 * 	余地。
	 * 	panic = exit
	 */
	fmt.Println("This is executed")
	/**
	 * 	panic: oh no.I can do no more.Goodbye.
	 *
	 *	goroutine 1 [running]:
	 *	main.main()
	 *	/Users/liukai/go/src/github.com/liukai6666/hello/errorHandle.go:125 +0x7e5
	 * 	exit status 2
	 */
	panic("oh no.I can do no more.Goodbye.")
	fmt.Println("This is not executed")

	/**
	 *	调用panic后，程序将停止执行，因此打印This is not executed的代码行根本没有机会执行。
	 *	在Go代码中，常常滥用下面的做法，这实际是说：朋友，咱们无路可走，只能让程序崩溃了。
	 *
	 * 	if err != nil {
	 *		panic(err)
	 * 	}
	 *
	 * 	在下面的情形下，使用panic可能是正确的选择
	 * 		1. 程序处于无法恢复的状态。这可能意味着无路可走了，或者再往下执行程序将带来更多的问题。在这种情况下，最佳的选择是让程序崩溃。
	 * 		2. 发生了无法处理的错误
	 */
}
