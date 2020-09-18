package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/**
	创建一个类似于map[string]interface{}的切片
	*/
	var s1 []map[string]interface{}

	/**
	使用make函数初始化map以开辟内存空间,
	*/
	m1 := make(map[string]interface{})

	/**
	为map进行赋值操作
	*/
	m1["name"] = "李白"
	m1["role"] = "打野"

	m2 := make(map[string]interface{})
	m2["name"] = "王昭君"
	m2["role"] = "中单"

	m3 := make(map[string]interface{})
	m3["name"] = "程咬金"
	m3["role"] = "上单"

	/**
	将map追加到切片中
	*/
	s1 = append(s1, m3, m2, m1)

	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("序列化出错,错误原因: ", err)
		return
	}

	/**
	查看序列化后的数据
	*/
	fmt.Println(string(data))
}
