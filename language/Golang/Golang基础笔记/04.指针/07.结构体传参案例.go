package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func addAge(s Student) {
	s.Age += 1
}

/**
在实际开发中,一般传递的结构体对象都是指针传递
*/
func addAgePointer(s *Student) {
	s.Age += 2
}

func main() {

	s1 := Student{"Jason Yin", 18}
	addAge(s1)
	fmt.Printf("姓名:[%s],年龄[%d]\n", s1.Name, s1.Age)
	/**
	结构体对象是值传递，但一般函数传递的都是结构体对象指针
	*/
	addAgePointer(&s1)
	fmt.Printf("姓名:[%s],年龄[%d]\n", s1.Name, s1.Age)
}
