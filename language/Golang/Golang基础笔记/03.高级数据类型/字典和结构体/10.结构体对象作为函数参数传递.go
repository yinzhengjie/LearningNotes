package main

import "fmt"

type Car struct {
	Name  string
	color string
	tyre  int8
}

func Demo(c Car) {
	c.color = "红色"
	fmt.Println("in Demo ...", c)
}

func main() {

	d1 := Car{"特斯拉", "蓝色", 4}
	fmt.Println("in main ...", d1)

	/*
		结构体对象作为函数参数,是值传递
	*/
	Demo(d1)

	fmt.Println("in main ...", d1)
}
