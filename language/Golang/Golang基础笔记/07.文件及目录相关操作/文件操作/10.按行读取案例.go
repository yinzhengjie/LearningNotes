package main

import (
	"bufio"
	"bytes"
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

	buf := make([]byte, 1024*4)
	f.Read(buf)                                     //一定要将数据读到切片中，否则执行下面的操作会读取不到数据哟~
	reader := bufio.NewReader(bytes.NewReader(buf)) //初始化一个阅读器
	line, _ := reader.ReadString('\n')              //为阅读器指定分隔符为换行符("\n"),即每次只读取一行.
	fmt.Println(line)
}
