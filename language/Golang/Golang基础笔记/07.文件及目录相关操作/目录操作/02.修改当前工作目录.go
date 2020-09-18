package main

import (
	"fmt"
	"os"
)

func main() {

	/**
	修改当前的工作目录，类似于linux操作系统中的cd命令。
	*/
	os.Chdir("E:\\yinzhengjie\\input")

	/**
	"."表示指定路径为当前路径,".."表示指定路径为当前路径的上级路径
	*/
	f, err := os.OpenFile("..", os.O_RDONLY, os.ModeDir)

	if err != nil {
		fmt.Println("目录打开失败: ", err)
		return
	}
	defer f.Close()

	files, err := f.Readdir(-100)
	if err != nil {
		fmt.Println("错误信息: ", err)
		return
	}

	for _, file := range files {
		fmt.Printf("文件名称是: %s,文件大小是:%d,是否是目录:%t\n", file.Name(), file.Size(), file.IsDir())
	}
}
