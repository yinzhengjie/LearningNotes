package main

import (
	"fmt"
	"net/rpc/jsonrpc"
)

func main() {
	/**
	首选是通过rpc.Dial拨号RPC服务

	温馨提示:
		默认数据传输过程中编码方式是gob,可以选择json,需要导入"net/rpc/jsonrpc"包。
	*/
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("链接服务器失败")
		return
	}
	defer conn.Close()

	var data string

	/**
	其中Call函数的签名如下所示:
		func (client *Client) Call(serviceMethod string, args interface{}, reply interface{}) error
	以下是对函数签名的相关参数进行补充说明:
		serviceMethod:
			用点号(.)链接的RPC服务名字和方法名字
		args:
			指定输入参数
		reply:
			指定输出参数接收的
	*/
	err = conn.Call("open_falcon.MonitorHosts", "Httpd", &data)
	if err != nil {
		fmt.Println("远程接口调用失败,错误原因: ", err)
		return
	}
	fmt.Println(data)
}
