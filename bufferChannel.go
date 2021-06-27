/**
 * 	使用缓冲通道
 * 		通常，通道收到消息后就可将其发送给接收者，但有时候可能没有接收者。在这种情况下，可使用缓冲通道。缓冲意味着可将数据存储在通道中，等接收者准备就绪再交给它
 * 	要创建缓冲通道，可向内置函数make传递另一个表示缓冲区长度的参数。
 * 		message := make(chan string, 2)
 * 		该代码创建一个可存储两条消息的缓冲通道。现在可在这个通道中添加两条消息了，虽然没有接收者。请注意，缓冲通道最多只能存储指定数量的消息，如果向它发送更多的消息将导致错误。
 */
package main

import (
	"fmt"
	"time"
)

func receiver(c chan string) {

	for msg := range c {
		fmt.Println(msg)
	}
}

/**
 * 	程序顺序执行、同步、并为使用go来做并行操作
 */
func main() {
	message := make(chan string, 2)
	message <- "hello"
	message <- "world"
	//	message <- "world"
	//关闭通道
	close(message)
	fmt.Println("Pushed two messages onto Channel with no receivers")
	time.Sleep(time.Second * 1)

	receiver(message)
}
