package main

import (
	"fmt"
)

func main() {
	m1 := map[string]rune{"first": '尹', "second": '正', "third": '杰'}

	//增加map键值
	fmt.Println("增加key之前:", m1)
	m1["test"] = 666666
	fmt.Println("增加key之后:", m1)

	//更新键值,
	m1["test"] = 88888888
	fmt.Println("更新key之后:", m1)

	//删除键值,Go语言种delete函数只是删除map种元素的作用
	delete(m1, "test")
	fmt.Println("删除key之后:", m1)

}
