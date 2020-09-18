package main

import (
	"fmt"
)

/**
定义函数，函数的返回值是函数类型
*/
func CallBack(a int) func(b int) int {
	return func(c int) int {
		fmt.Println("调用了CallBack这个回调函数...")
		return a + c
	}
}

type BigData struct {
	Name string
}

func (this *BigData) Init() {
	this.Name = "Hadoop"
}

func (this *BigData) PrinfInfo() {
	fmt.Printf("%v是大数据生态圈的基石。\n", this.Name)
}

func (this BigData) SetInfoValue() {
	fmt.Printf("SetInfoValue : %p,%v\n", &this, this)
}

func (this *BigData) SetInfoPointer() {
	fmt.Printf("SetInfoPointer : %p,%v\n", this, this)
}

func main() {

	/**
	  调用回调函数的返回值为函数类型
	*/
	result := CallBack(10)
	fmt.Printf("result的类型是:[%T],result的值是:[%v]\n", result, result)
	res1 := result(20) //我们对返回的函数再次进行调用
	fmt.Printf("res1的类型是:[%T],res1的值是:[%d]\n\n", res1, res1)

	var hadoop BigData
	hadoop.Init()            //调用hadoop的初始化方法
	info := hadoop.PrinfInfo //我们可以声明一个函数变量info,我们称之为方法表达式

	/**
	  我们对info函数变量进行调用，这样可以起到隐藏调用者hadoop对象的效果哟~(类似于回调函数的调用效果)
	  方法值可以隐藏调用者,我们称为隐式调用。
	*/
	info()

	/**
	　　方法表达式可以显示调用调用，必须传递方法调用者对象，在实际开发中很少使用这种方式，我们了解即可。
	*/

	elk := BigData{"Elastic Stack"}
	fmt.Printf("main:%p,%v\n\n", &elk, elk)
	s1 := (*BigData).SetInfoPointer
	s1(&elk) //显示把接收者传递过去
	s2 := (BigData).SetInfoValue
	s2(elk) //显示把接收者传递过去
}
