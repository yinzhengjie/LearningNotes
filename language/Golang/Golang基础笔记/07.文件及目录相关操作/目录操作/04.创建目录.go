package main

import (
	"os"
)

func main() {

	/**
	切换工作路径
	*/
	os.Chdir("E:\\yinzhengjie\\input")

	/**
	切换到指定的工作路径后创建bigdata目录
	*/
	os.Mkdir("bigdata", 0755)

}
