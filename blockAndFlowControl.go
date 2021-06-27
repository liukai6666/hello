/**
 * 		Goroutine是Go语言提供的一种并发编程方式，速度缓慢的网络调用或函数会阻塞程序的执行，而Goroutine能够让您对此进行管理。在并发编程中，通常应避免阻塞式操作，但有时需要让代码处于
 * 	阻塞状态。例如，需要在后台运行的程序必须阻塞，这样才不会退出
 * 		Goroutine会立即返回（非阻塞），因此要让进程处于阻塞状态，必须采用一些流程控制技巧。例如，从通道接收并打印消息的程序需要阻塞，以免终止。
 * 		给通道指定消息接收者是一个阻塞操作，因为它将阻止函数返回，直到收到一条消息为止
 * 		Goroutine是非阻塞的，因此如果程序要阻塞，以接收大量的消息或不断地重复某个过程，必须使用其他流程控制技术
 */
package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	t := time.NewTicker(5 * time.Second)

	/**
	 * 	死循环、每5秒写入一个ping
	 *
	 */
	for {
		c <- "ping"
		<-t.C
	}
}

func main() {
	messages := make(chan string)
	go pinger(messages)

	/**
	 *	死循环、不断接收message管道的消息，并打印
	 */
	for {
		msg := <-messages
		fmt.Println(msg)
	}

}
