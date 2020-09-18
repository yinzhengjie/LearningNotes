package main

import (
	"encoding/json"
	"fmt"
)

/**
定义需要结构体
*/
type Teacher struct {
	Name    string
	ID      int
	Age     int
	Address string
}

func main() {
	s1 := Teacher{
		Name:    "Jason Yin",
		ID:      001,
		Age:     18,
		Address: "北京",
	}

	/**
	使用“encoding/json”包的Marshal函数进行序列化操作,其函数签名如下所示:
		func Marshal(v interface{}) ([]byte, error)
	以下是对Marshal函数参数相关说明:
		v:
			该参数是空接口类型。意味着任何数据类型(int、float、map,结构体等)都可以使用该函数进行序列化。
		返回值:
			很明显返回值是字节切片和错误信息
	*/
	//data, err := json.Marshal(&s1)

	/**
	Go语言标准库的"encoding/json"包还提供了另外一个方法：MarshalIndent。
	该方法的作用与Marshall作用相同，只是可以通过方法参数，设置前缀、缩进等，对Json多了一些格式处理，打印出来比较好看。
	*/
	data, err := json.MarshalIndent(s1, "\t", "")
	if err != nil {
		fmt.Println("序列化出错,错误原因: ", err)
		return
	}

	/**
	查看序列化后的json字符串
	*/
	fmt.Println("序列化之后的数据为: ", string(data))
}
