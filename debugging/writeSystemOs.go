package main

import (
	"log"
)

func main() {

	for i := 1; i <= 5; i++ {
		log.Printf("Log iteration %d", i)
	}

	/**
	 * 	通过使用重定向功能，可将输出到重定向到文件
	 * 	go run writeSystemOs.go > example04.log 2>&1
	 */
}
