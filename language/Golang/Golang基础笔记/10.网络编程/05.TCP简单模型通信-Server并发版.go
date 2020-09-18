package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	//函数调用完毕，自动关闭conn
	defer conn.Close()

	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, " conncet sucessful")

	buf := make([]byte, 2048)

	for {
		//读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
		fmt.Println("len = ", len(string(buf[:n])))

		//if "exit" == string(buf[:n-1]) { 	// nc测试，发送时，只有 \n
		if "exit" == string(buf[:n-2]) { // 自己写的客户端测试, 发送时，多了2个字符, "\r\n"
			fmt.Println(addr, " exit")
			return
		}

		//把数据转换为大写，再给用户发送
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

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

	//接收多个用户
	for {
		/**
		  等待客户端连接请求
		*/
		conn, err := socket.Accept()
		if err != nil {
			fmt.Println("建立链接失败,错误原因: ", err)
			return
		}

		//处理用户请求, 新建一个go程
		go HandleConn(conn)
	}
}
