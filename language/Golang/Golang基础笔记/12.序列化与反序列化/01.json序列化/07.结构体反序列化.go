package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name    string
	Age     int
	Address string
}

func main() {
	/**
	以Json数据为例,我们接下来要对该数据进行反序列化操作。
	*/
	p1 := `{"Name":"Jason Yin","Age":18,"Address":"北京"}`

	var s1 People
	fmt.Printf("反序列化之前: \n\ts1 = %v \n\ts1.Name = %s\n\n", s1, s1.Name)

	/**
	使用encoding/json包中的Unmarshal()函数进行反序列化操作,其函数签名如下:
		func Unmarshal(data []byte, v interface{}) error
	以下是对函数签名的参数说明:
		data:
			待解析的json编码字符串
		v:
			解析后传出的结果,即用来可以容纳待解析的json数据容器.
	*/
	err := json.Unmarshal([]byte(p1), &s1)
	if err != nil {
		fmt.Println("反序列化失败: ", err)
		return
	}

	/**
	查看反序列化后的结果
	*/
	fmt.Printf("反序列化之后: \n\ts1 = %v \n\ts1.Name = %s\n", s1, s1.Name)

}
