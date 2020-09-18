package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args //获取命令行参数， 并判断输入是否合法

	if args == nil || len(args) != 3 {
		fmt.Println("useage : xxx srcFile dstFile")
		return
	}

	srcPath := args[1] //获取源文件路径
	dstPath := args[2] //获取目标文件路径

	fmt.Printf("srcPath = %s, dstPath = %s\n", srcPath, dstPath)

	if srcPath == dstPath {
		fmt.Println("error：源文件名与目的文件名相同")
		return
	}

	srcFile, err1 := os.Open(srcPath) // 打开源文件
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	dstFile, err2 := os.Create(dstPath) //创建目标文件
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	buf := make([]byte, 1024) //切片缓冲区
	for {
		//从源文件读取内容，n为读取文件内容的长度
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}

		if n == 0 {
			fmt.Println("文件处理完毕")
			break
		}

		//切片截取
		tmp := buf[:n]

		//把读取的内容写入到目的文件
		dstFile.Write(tmp)
	}

	//关闭文件
	srcFile.Close()
	dstFile.Close()
}
