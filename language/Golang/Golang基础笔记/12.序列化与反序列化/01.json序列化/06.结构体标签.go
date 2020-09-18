package main

import (
	"encoding/json"
	"fmt"
)

/**
结构体的字段除了名字和类型外，还可以有一个可选的标签（tag）,它是一个附属于字段的字符串，可以是文档或其他的重要标记。
比如在我们解析json或生成json文件时，常用到encoding/json包，它提供一些默认标签。
定义结构体时，可以通过这些默认标签来设定结构体成员变量，使之在序列化后得到特殊的输出。
*/
type Student struct {
	/**
	“-”标签:
		作用是不进行序列化，效果和将结构体字段首字母写成小写一样。
	*/
	Name string `json:"-"`

	/**
	string标签:
		这样生成的json对象中，ID的类型转换为字符串
	*/
	ID int `json:"id,string"`

	/**
	omitempty标签：
		可以在序列化的时候忽略0值或者空值；
	*/
	Age int `json:"AGE,omitempty"`

	/**
	可以将字段名称进行重命名操作:
		比如下面的案例就是将"Address"字段重命名为"HomeAddress"哟~
	*/
	Address string `json:"HomeAddress"`

	/**
	由于该字段首字母是小写，因此该字段不参与序列化哟~
	*/
	score int
	Hobby string
}

func main() {
	s1 := Student{
		Name: "Jason Yin",
		ID:   001,
		//Age:     18,
		Address: "北京",
		score:   100,
		Hobby:   "中国象棋",
	}

	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("序列化出错,错误原因: ", err)
		return
	}
	fmt.Println("序列化结果: ", string(data))
}
