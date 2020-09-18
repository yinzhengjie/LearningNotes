package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {

	/**
	使用Listen函数创建监听socket,其函数签名如下:
		func Listen(network, address string) (Listener, error)
	以下是对函数签名的参数说明:
		network:
			指定服务端socket的协议,如tcp/udp,注意是小写字母哟~
		address:
			指定服务端监听的IP地址和端口号,如果不指定地址默认监听当前服务器所有IP地址哟~
	*/
	socket, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("开启监听失败,错误原因: ", err)
		return
	}
	defer socket.Close()
	fmt.Println("开启监听...")

	for {
		/**
		等待客户端连接请求
		*/
		conn, err := socket.Accept()
		if err != nil {
			fmt.Println("建立链接失败,错误原因: ", err)
			return
		}
		defer conn.Close()
		fmt.Println("建立链接成功,客户端地址是: ", conn.RemoteAddr())

		/**
		分两次接收客户端数据:
			第一次最终接收数据的长度;
			第二次根据第一次接受的长度，创建容量大小;
		*/
		tmp := make([]byte, 2)

		conn.Read(tmp)
		dataLength, err := strconv.Atoi(string(tmp)) //把字节切片转换成整型
		if err != nil {
			fmt.Println("获取数据长度失败: ", err)
			return
		}
		fmt.Println("获取到的数据长度是: ", dataLength)

		conn.Write([]byte("已获取到数据长度"))

		/**
		开始读取数据
		*/
		buf := make([]byte, dataLength)
		conn.Read(buf)
		fmt.Printf("读取到客户端的数据为: %s\n", string(buf))

		/**
		发送数据给客户端
		*/
		data := "Blog地址:[https://www.cnblogs.com/yinzhengjie/]"
		conn.Write([]byte(data))
	}
}
