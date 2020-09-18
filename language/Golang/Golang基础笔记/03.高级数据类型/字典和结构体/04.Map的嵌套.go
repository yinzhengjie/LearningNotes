package main

import (
	"fmt"
)

func main() {

	m1 := map[string]rune{"first": '尹', "second": '正', "third": '杰'}

	/*
		定义一个嵌套数据类型
	*/
	m2 := make(map[string]map[string]int32)

	//我们可以为嵌套类型赋值
	m2["name"] = m1
	fmt.Println("m1的数据为:", m1)
	fmt.Println("m2的数据为:", m2)

}
