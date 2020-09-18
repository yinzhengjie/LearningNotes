package main

import (
	"fmt"
	"net"
)

func main() {
	/**
	使用Dial函数链接服务端,其函数签名如下所示:
		func Dial(network, address string) (Conn, error)
	以下是对函数签名的各参数说明:
		network:
			指定客户端socket的协议,如tcp/udp,该协议应该和需要链接服务端的协议一致哟~
		address:
			指定客户端需要链接服务端的socket信息,即指定服务端的IP地址和端口
	*/
	conn, err := net.Dial("udp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("连接服务端出错,错误原因: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("与服务端连接建立成功...")

	/**
	给服务端发送数据
	*/
	conn.Write([]byte("Hi,My name is Jason Yin."))

	/**
	读取服务端返回的数据
	*/
	tmp := make([]byte, 1024)
	n, _ := conn.Read(tmp)
	fmt.Println("获取到服务器返回的数据为: ", string(tmp[:n]))
}
