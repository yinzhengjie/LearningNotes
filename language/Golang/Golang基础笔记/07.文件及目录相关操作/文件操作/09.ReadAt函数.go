package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("E:\\yinzhengjie\\input\\flume.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败: ", err)
		return
	}
	defer f.Close()

	temp := make([]byte, 1024*4)

	f.ReadAt(temp, 5) //带偏移量的获取数据，是从文件头开始的,当然我们使用Seek函数也能实现该功能

	fmt.Println(string(temp))
}
