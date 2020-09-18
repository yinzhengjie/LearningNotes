package main

import (
	"fmt"
)

type Father struct {
	Name string
	Age  int
}

func (f *Father) Init() {
	f.Name = "成龙"
	f.Age = 66
}

//定义父类的Eat成员方法
func (f *Father) Eat() {
	fmt.Println("Jackie Chan is eating...")
}

//重写父类的Eat成员方法
func (s *Son) Eat() {
	fmt.Println("FangZuming is eating...")
}

//我们让Son类继承Ｆather父类
type Son struct {
	Father //匿名组合能够继承父类的属性和方法
	Score  int
}

func main() {
	var s Son
	s.Init()
	fmt.Printf("%+v\n", s)
	s.Eat()
	s.Name = "房祖名"
	s.Age = 38
	s.Score = 100
	fmt.Printf("%+v\n", s)
}
