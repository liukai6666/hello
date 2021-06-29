/**
 * 	如何将日志消息写入文件
 */
package main

import (
	"errors"
	"log"
	"os"
)

func main() {

	/**
	 * 2021/06/29 12:58:26 We only just started and we are crashing
	 * exit status 1
	 */
	var errFatal = errors.New("We only just started and we are crashing")
	log.Fatal(errFatal)

	/**
	 *  	日志可用来了解程序的执行情况，它对调试能起到一定作用，但是，查看日志对了解发生的意外事件更有帮助。这意味着需要将日志写入文件，以便能够访问它们。要将日志写入文件，可使用Go语言
	 *  本身提供的功能。要将日志写入文件，只需将log包这样做即可
	 */
	f, err := os.OpenFile("example03.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	//最后一行让log包使用文件example03.log来记录应用程序的日志，
	log.SetOutput(f)

}
