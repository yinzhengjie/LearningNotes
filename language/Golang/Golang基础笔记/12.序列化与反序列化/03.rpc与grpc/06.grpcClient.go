package main

import (
	"google.golang.org/grpc"
	"context"
	"fmt"
	"yinzhengjie/pb"
)

func main()  {

	//和grpc服务端建立连接
	grpcCnn ,err := grpc.Dial("127.0.0.1:8888",grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer grpcCnn.Close()

	//得到一个客户端对象
	client :=pb.NewHelloClient(grpcCnn)

	var s pb.Student
	s.Name = "Jason Yin"
	s.Age = 20

	resp,err := client.World(context.TODO(),&s)
	fmt.Println(resp,err)
}
