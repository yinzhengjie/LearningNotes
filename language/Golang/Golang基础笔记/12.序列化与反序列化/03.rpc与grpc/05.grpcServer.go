package main

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"yinzhengjie/pb"
)


//定义一个结构体,继承自HelloServer接口(该接口是我们通过protobuf代码生成的)
type HelloService struct {}

func (HelloService)World(ctx context.Context, req*pb.Student) (*pb.Student, error){
	req.Name += " nihao"
	req.Age += 10
	return req,nil
}


func main()  {
	//先获取grpc对象
	grpcServer := grpc.NewServer()

	//注册服务
	pb.RegisterHelloServer(grpcServer,new(HelloService))

	//开启监听
	lis,err := net.Listen("tcp",":8888")
	if err != nil {
		return
	}
	defer lis.Close()

	//先获取grpc服务端对象
	grpcServer.Serve(lis)
}
