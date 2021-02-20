package main

import (
	"fmt"
	"net"
)

/**
tcp server端
*/
func processConn(conn net.Conn) {
	//与客户端通信
	var tmp [128]byte
	read, err := conn.Read(tmp[:])
	if err != nil {
		fmt.Println("读取失败err:", err)
		return
	}
	fmt.Println(string(tmp[:read]))
}

func main() {
	//本地端口启动服务
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("启动server失败err:", err)
		return
	}
	//等待别人跟server建立连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("建立连接失败err:", err)
			return
		}
		go processConn(conn)
	}
}
