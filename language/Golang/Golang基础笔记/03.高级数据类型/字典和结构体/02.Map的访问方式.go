package main

import (
	"fmt"
)

func main() {
	m1 := map[string]rune{"first": '尹', "second": '正', "third": '杰'}

	//第一种访问方式,可以通过key值访问
	fmt.Println("===== 第一种访问方式 =====")
	fmt.Println(m1["first"])

	//第二种访问方式,可以通过变量名访问所有数据
	fmt.Println("===== 第二种访问方式 =====")
	fmt.Println(m1)

	//第三种访问方式,同时拿到key和value
	fmt.Println("===== 第三种访问方式 =====")
	for key, value := range m1 {
		fmt.Println("key值是:", key, ",value值是:", value)
	}

	//第四种访问方式,只拿到key,基于key去范围其对应的value
	fmt.Println("===== 第四种访问方式 =====")
	for key := range m1 {
		fmt.Println("key值是:", key, ",value值是:", m1[key])
	}

	//第五种访问方式,判断一个map是否有key,基于返回的bool值做相应的操作
	fmt.Println("===== 第五种访问方式 =====")
	value, flag := m1["first"]
	if flag {
		fmt.Println("key的值为:", value)
	}
}
