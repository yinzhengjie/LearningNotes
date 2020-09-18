package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

type Student struct {
	Person //通过匿名组合的方式嵌入了Person的属性。
	Score  float64
}

type Teacher struct {
	Person //通过匿名组合的方式嵌入了Person的属性。
	Course string
}

type Schoolmaster struct {
	Person   //通过匿名组合的方式嵌入了Person的属性。
	CarBrand string
}

func main() {
	/**
	第一种初始化方式:先定义后赋值
	*/
	s1 := Student{}
	s1.Name = "Jason Yin"
	fmt.Println(s1)
	fmt.Printf("%+v\n\n", s1) //"+v表示打印结构体的各个字段"

	/**
	第二种初始化方式:直接初始化
	*/
	s2 := Teacher{Person{"尹正杰", 18, "boy"}, "Go并发编程"}
	fmt.Println(s2)
	fmt.Printf("%+v\n\n", s2)

	/**
	第三种赋值方式:初始化赋值部分字段
	*/
	s3 := Schoolmaster{CarBrand: "丰田", Person: Person{Name: "JasonYin最强王者"}}
	fmt.Println(s3)
	fmt.Printf("%+v\n", s3)
}
