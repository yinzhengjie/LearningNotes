package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Zabbix struct{}

/**
定义成员方法:
	第一个参数是传入参数.
	第二个参数必须是传出参数(引用类型).

Go语言的RPC规则：
	方法只能有两个可序列化的参数，其中第二个参数是指针类型，并且返回一个error类型，同时必须是公开的方法。

温馨提示:
	当调用远程函数之后，如果返回的错误不为空，那么传出参数为空。
*/
func (Zabbix) MonitorHosts(name string, response *string) error {
	*response = name + "主机监控中..."
	return nil
}

func main() {
	/**
	进程间交互有很多种，比如基于信号,共享内存,管道,套接字等方式.

	1>.rpc基于是TCP的,因此我们需要先开启监听端口
	*/
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("开启监听器失败,错误原因: ", err)
		return
	}
	defer listener.Close()
	fmt.Println("服务启动成功...")

	/**
	2>.接受链接，即接受传输的数据
	*/
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("建立链接失败...")
		return
	}
	defer conn.Close()
	fmt.Println("建立连接: ", conn.RemoteAddr())
	/**
	3>.注册rpc服务，维护一个hash表，key值是服务名称，value值是服务的地址。服务器有很多函数，希望被调用的函数需要注册到RPC上。
	以下是RegisterName的函数签名:
		func RegisterName(name string, rcvr interface{}) error
	以下是对函数签名相关参数的说明:
		name:
			指的是服务名称。
		rcvr:
			指的是结构体对象(这个结构体必须含有成员方法)。
	*/
	rpc.RegisterName("zabbix", new(Zabbix))

	/**
	4>.链接的处理交给RCP框架处理,即rpc调用,并返回执行后的数据,其工作原理大致分为以下3个步骤:
		(1)read,获取服务名称和方法名，获取请求数据;
		(2)调用对应服务里面的方法，获取传出数据;
		(3)write,把数据返回给client;
	*/
	rpc.ServeConn(conn)
}
