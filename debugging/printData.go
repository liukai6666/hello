/**
 * 	提示玩家输入，并根据输入判断玩家能否获得奖励。如果玩家猜对了名字，就将获得奖励；否则将显示一条消息，指出他没有获得奖励
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//读取从系统输入
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Guess the name of my pet to win a prize：")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	fmt.Println("[DEBUG]text is:", text)
	if text == "john" {
		fmt.Println("You won! you win chocolate!")
	} else {
		fmt.Println("You didn't win, Better luck next time")
		main()
	}

}
