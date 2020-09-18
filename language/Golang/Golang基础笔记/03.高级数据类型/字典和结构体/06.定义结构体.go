package main

import (
	"fmt"
)

/*
	我们定义一个变量后不使用运行会报错,但定义结构体类型后可以不使用,在执行时并不会报错,因为你不对结构体进行初始化操作它并不占用内存空间.
*/
type Teacher struct {
	Name string
	Age  uint8
}

/*
	温馨提示:
		Go语言中结构体的字段首字母如果小写的话不能跨包访问,如果首字母大写则可以跨包访问.

	如下所示,我们可以自定义一个全局类型的Student结构体类型
*/
type Student struct {
	Name   string
	Age    uint8
	Gender rune
}

func main() {
	//初始化一个结构体
	var s1 Student
	s1.Name = "尹正杰"
	s1.Age = 18
	fmt.Println("s1 = ", s1)

	//我们也可以在声明变量的时候定义结构体类型,只不过我们实在main函数声明的,因此是一个局部变量
	var s2 struct {
		Name  string
		Score int
	}
	s2.Name = "Jason Yin"
	s2.Score = 750
	fmt.Println("s2 = ", s2)

	//我们可以直接进行初始化赋值
	s3 := Student{Name: "Yinzhengjie", Age: 27, Gender: '男'}
	s4 := Student{"尹正杰", 27, '男'} //可以省略字段名称但顺序一定不能错
	fmt.Println("s3 = ", s3)
	fmt.Println("s4 = ", s4)
}
