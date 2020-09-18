package main

import (
	"fmt"
	"net"
)

func main() {
	/**
	创建监听的地址，并且指定udp协议
	*/
	udp_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("获取监听地址失败,错误原因: ", err)
		return
	}

	/**
	创建数据通信socket
	*/
	conn, err := net.ListenUDP("udp", udp_addr)
	if err != nil {
		fmt.Println("开启UDP监听失败,错误原因: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("开启监听...")

	buf := make([]byte, 1024)

	/**
	通过ReadFromUDP可以读取数据，可以返回如下三个参数:
		dataLength:
			数据的长度
		raddr:
			远程的客户端地址
		err:
			错误信息
	*/
	dataLength, raddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("获取客户端传递数据失败,错误原因: ", err)
		return
	}
	fmt.Println("获取到客户端的数据为: ", string(buf[:dataLength]))

	/**
	写回数据
	*/
	conn.WriteToUDP([]byte("服务端已经收到数据啦~"), raddr)
}
