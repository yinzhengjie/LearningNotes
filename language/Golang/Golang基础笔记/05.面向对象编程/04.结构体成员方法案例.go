package main

import (
	"fmt"
)

//定义一个结构体
type Lecturer struct {
	Name string
	Age  uint8
}

//我们为Lecturer结构体封装Init成员方法
func (l *Lecturer) Init() {
	l.Name = "Jason Yin"
	l.Age = 20
}

/**
我为Lecturer结构体起一个别名
我们可以为Instructor类型添加成员方法,
通过别名和成员方法为原有类型赋值新的操作
*/
type Instructor Lecturer

/**
温馨提示:
	(1)我们为一个结构体创建成员方法时，如果成员方法有接收者,需要考虑以下两种情况:
		1>.如果这个接收者是对象的时候，是值传递;
		2>.如果这个接收者是对象指针，是引用传递;
	(2)只要函数接收者不同,哪怕函数名称相同，也不算同一个函数哟;
	(3)不管接收者变量名称是否相同,只要类型一致(包括对象和对象指针),那么我们就认为接收者是相同的,这时候不允许出现相同名称函数;
	(4)给指针添加方法的时候，不允许给指针类型添加操作(因为Go语言中指针类型是只读的);
*/
func (i *Instructor) Init() {
	i.Name = "尹正杰"
	i.Age = 18
}

func main() {

	var (
		l Lecturer
		i Instructor
	)

	//可以使用对象调用成员方法
	i.Init()
	fmt.Printf("%+v\n\n", i)

	//可以用对象指针调用成员方法
	(&l).Init()
	fmt.Printf("%+v\n", l)
}
