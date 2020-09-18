package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var s1 map[string]interface{}

	/**
	使用make函数初始化map以开辟内存空间
	*/
	s1 = make(map[string]interface{})

	/**
	map赋值操作
	*/
	s1["name"] = "Jason Yin"
	s1["age"] = 20
	s1["address"] = [2]string{"北京", "陕西"}

	/**
	将map使用Marshal()函数进行序列化
	*/
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("Marshal err: ", err)
		return
	}

	/**
	查看序列化后的json字符串
	*/
	fmt.Println("序列化之后的数据为: ", string(data))

}
