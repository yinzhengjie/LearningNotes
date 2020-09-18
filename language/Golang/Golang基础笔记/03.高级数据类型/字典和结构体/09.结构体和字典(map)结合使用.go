package main

import (
	"fmt"
)

type Server struct {
	brand       string //品牌
	CpuCore     uint8  //核心数
	MemorySpace int    //内存空间,单位为GB
	DiskSpace   int    //磁盘空间,单位为GB
}

func main() {
	//使用自动推导类型直接进行初始化赋值
	s1 := Server{"戴尔", 32, 128, 122880}

	//定义一个字典
	m1 := make(map[string]Server)

	//结构体对象作为map的value
	m1["亦庄机房"] = s1
	fmt.Println(m1["亦庄机房"].brand)
	fmt.Println(m1)

	/*
	   如果map的value是结构体类型,那么不能直接通过map修改结构体中的数据,这个时候结构体的内容是只读的.
	   如果非要修改map的value是结构体类型,需要先对结构体进行修改,然后再重新赋值
	*/
	//m1["亦庄机房"].brand = "浪潮" //这种写法是错误的
	s1.brand = "浪潮" //先修改s1结构体的brand字段
	m1["亦庄机房"] = s1 //再对同一个key进行赋值,完成更新操作
	fmt.Println(m1["亦庄机房"].brand)
	fmt.Println(m1)
}
