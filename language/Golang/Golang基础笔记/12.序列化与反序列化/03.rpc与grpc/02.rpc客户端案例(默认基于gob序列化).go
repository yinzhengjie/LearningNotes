package main

import (
	"fmt"
	"net"
	"net/rpc"
)

func main() {
	/**
	1>.首选是通过rpc.Dial拨号RPC服务

	温馨提示:
		默认数据传输过程中编码方式是gob,可以选择json
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("链接服务器失败")
		return
	}
	defer conn.Close()

	/**
	2>.把conn和rpc进行绑定
	*/
	client := rpc.NewClient(conn)

	/**
	3>.然后通过client.Call调用具体的RPC方法。其中Call函数的签名如下所示:
			func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error
		以下是对函数签名的相关参数进行补充说明:
			serviceMethod:
				用点号(.)链接的RPC服务名字和方法名字
			args:
				指定输入参数
			reply:
				指定输出参数接收的
	*/
	var data string
	err = client.Call("zabbix.MonitorHosts", "Nginx", &data)
	if err != nil {
		fmt.Println("远程接口调用失败,错误原因: ", err)
		return
	}
	fmt.Println(data)
}
