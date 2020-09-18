package main

import (
	"fmt"
	"os"
)

func main() {

	/**
	获取当前的工作路径
	*/
	pwd, _ := os.Getwd()
	fmt.Println(pwd)

	/**
	修改当前的工作目录，类似于linux操作系统中的cd命令。
	*/
	os.Chdir("E:\\yinzhengjie\\input")

	pwd, _ = os.Getwd()
	fmt.Println(pwd)
}
