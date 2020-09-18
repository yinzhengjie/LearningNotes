package main

import (
	"fmt"
)

/*
	温馨提示:
		(1)结构体中的字段一旦定义就不会发生改变;
		(2)结构体没有增加或者删除字段的操作;
		(3)在main函数中调用其它文件的成员,在编译运行的时候,必须手动指定编译的文件;

*/
type Person struct {
	Name string
	Age  uint8
}

func main() {

	//对自定义的结构体进行初始化赋值
	s := Person{"尹正杰", 18}

	fmt.Println(s)

	//我们可以单独访问结构体的元素
	fmt.Printf("姓名:%s,年龄:%d\n", s.Name, s.Age)
}
