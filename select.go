/**
 * 		假设有多个Goroutine，而程序将根据最先返回的Goroutine执行相应的操作，此时可使用select语句，select语句类似于switch语句，它为通道创建一系列接收者，并执行最先收到消息的
 * 	接收者
 * 		select语句看起来和switch很像
 * 		channel1 := make(chan string)
 * 		channel2 := make(chan string)
 *
 * 		select {
 *			case msg1 := <-channel1:fmt.Println("received",msg1)
 * 			case msg2 := <-channel2:fmt.Println("received",msg2)
 *	 	}
 * 		如果从通道channel1那里收到了消息，将执行第一条case语句；如果从通道channel2那里收到了消息，将执行第二条case语句，取决于消息到达的时间，哪条消息最先到达决定了将执行哪条case
 * 	语句。通常，接下来收到的其他消息将被丢弃。收到一条消息后，select语句将不再堵塞
 */
package main

import (
	"fmt"
	"time"
)

func ping1(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "ping on channel"
}

func ping2(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "ping on channel2"
}

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go ping1(channel1)

	go ping2(channel2)

	/**
	 *	收到来自最先返回通道的消息后，执行相应的case语句---向终端打印一条消息。
	 * 	整个select语句就此结束，不再堵塞进程，因此程序退出
	 * 	case <-time.After(500 * time.Microsecond):
	 * 	指定超时时间
	 */
	select {
	case msg1 := <-channel1:
		fmt.Println("received", msg1)
	case msg2 := <-channel2:
		fmt.Println("received", msg2)
	case <-time.After(500 * time.Microsecond):
		fmt.Println("no message received. giving up.")
	}
}
