package main

import (
	"fmt"
	"net"
	"strconv"
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
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("连接服务端出错,错误原因: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("与服务端连接建立成功...")

	/**
	定义需要发送的数据,第一次给服务端发送要发的长度
	*/
	data := []byte("服务端,请问博客地址的URL是多少呢?")
	lenData := len(data)

	/**
	给服务端发送数据的长度
	*/
	conn.Write([]byte(strconv.Itoa(lenData)))

	/**
	获取服务器的应答
	*/
	var buf = make([]byte, 1024)
	conn.Read(buf)
	fmt.Printf("从服务端获取到的数据为:%s\n", string(buf))

	/**
	第二次给服务器发送数据
	*/
	conn.Write(data)
	conn.Read(buf)
	fmt.Printf("获取到的数据为:%s\n", string(buf))
}
