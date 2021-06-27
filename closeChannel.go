/**
 * 		在已知需要停止执行的时间的情况下，使用超时时间是不错的选择，但在有些情况下，不确定select语句该在何时返回，因此不能使用定时器。在这种情况下，可使用退出管道。
 * 	这种技术并非语言规范的组成部分，但可通过向通道发送消息来理解退出阻塞的select语句。
 * 		来看这样一种情形：程序需要使用select语句实现无限制地阻塞，但同时要求能够随时返回。通过在select语句中添加一个退出通道，可向退出通道发送消息来结束该语句，从而停止阻塞。可
 * 	将退出通道视为阻塞式select语句的开关。对于退出通道，可随便命名，但通常将其命名为stop或quit。在下面的示例中，在for循环中使用了一条select语句，这意味着它将无限制地阻塞，并不断
 * 	地接收消息。通过向通道stop发送消息，可让select语句停止阻塞：从for循环中返回，并继续往下执行。
 * 		messages := make(chan string)
 * 		stop := make(chan bool)
 *
 * 		for {
 * 			select {
 * 			case <-stop:
 * 				return
 * 			case msg:=<-messages:
 * 				fmt.Println(msg)
 *	 		}
 *	 	}
 * 		在应用程序的某部分向通道发送消息，并要在未来的某个位置时点终止时，这种技术是很有效的。
 * 		为了提供这样的示例，我们在Goroutine中创建一个函数，它每隔1s向通道发送一条消息。
 * 		func sender(c chan string) {
 * 			t := time.NewTriker(1 * time.Second)
 * 			for {
 * 				c <- "I'm sending a message"
 * 				<-t.C
 *	 		}
 *		}
 * 		messages := make(chan string)
 * 		go sender(messages)
 * 		通过在for循环中使用select语句，可在收到消息后立即打开它。由于这是一个阻塞操作，因此将不断打印消息，直到您手动终止这个过程。
 * 		for {
 *			select {
 *			case msg := <-messages:
 * 				fmt.Println(msg)
 *			}
 *	 	}
 */
package main

import (
	"fmt"
	"time"
)

/**
 * 	死循环，每隔1秒钟向管道中发送1条消息 I'm sending a message
 */
func sender(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "I'm sending a message"
		<-t.C
	}
}

func main() {
	//创建发送消息管道
	messages := make(chan string)
	//创建停止消息管道
	stop := make(chan bool)
	//并行执行sender
	go sender(messages)

	//并行执行、休眠两秒、两秒到时，打印Time's up! 给停止消息的管道赋值true
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Time's up")
		stop <- true
	}()
	/**
	 * 	死循环、 如果检测到停止信号，return 、 如果接收到messages管道发送过来的消息，则打印消息
	 */
	for {
		select {
		case msg := <-messages:
			fmt.Println(msg)
		case <-stop:
			return
		}
	}
}
