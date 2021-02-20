package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/**
tcp client端
*/
func main() {
	//与server端建立连接
	dial, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("client建立连接失败err:", err)
		return
	}
	//发送数1据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("清输入发送内容")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		fmt.Scanln(&msg)
		if msg == "exit" {
			break
		}
		dial.Write([]byte(msg))
	}

	defer dial.Close()
}
