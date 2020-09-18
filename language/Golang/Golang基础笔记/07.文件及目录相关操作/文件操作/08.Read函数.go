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
	/**
	切片在使用前要申请内存空间,如果不知道文件的大小，尽量多给点空间最好是4K的倍数
	如果在读取大文件的情况下，我们应该循环读取，当然，我的笔记里有读取大文件的实战案例。
	*/
	temp := make([]byte, 1024*4)

	f.Read(temp)

	fmt.Println(string(temp))
}
