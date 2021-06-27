/**
 * 	将通道用作函数参数
 * 		您在前面的示例中了解过，可将通道作为参数传递给函数，并在函数中向通道发送消息。要进一步指定在函数中如何使用传入的通道，可在传递通道时将其指定为只读、只写、或读写的。指定
 * 	通道是只读、只写、读写的语法差别不大。
 * 		<- 位于chan左边时，表示通道在函数内是只读的
 * 		<- 位于chan右边时，表示通道在函数内是只写的
 * 		没有指定<-时，表示通道是可读可写的
 * 		通过指定通道访问权限，有助于确保通道中数据的完整性，还可指定程序的哪部分可向通道发送数据或接收来自通道的数据
 *
 */
package main

import (
	"fmt"
)

func channelReader(messages <-chan string) {
	msg := <-messages
	fmt.Println(msg)
}

func channelWriter(messages chan<- string) {
	messages <- "hello world"
}

func channelReaderAndWriter(messages chan string) {
	msg := <-messages
	fmt.Println(msg)
	messages <- "hi,girl"
}
