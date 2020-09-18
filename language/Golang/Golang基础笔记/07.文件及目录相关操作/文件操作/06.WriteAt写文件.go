package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("E:\\yinzhengjie\\input\\kafka.txt")
	if err != nil {
		fmt.Println(errors.New("报错提示:" + err.Error()))
		return
	} else {
		fmt.Println("文件创建成功....")
	}

	/**
	文件成功创建之后,一直处于打开状态，因此我们使用完之后一定要关闭文件，当然关闭文件之前一定要确保为文件已经创建成功。
	*/
	defer f.Close()

	/**
	WriteAt是带偏移量的写入数据，偏移量是从文件起始位置开始,WriteAt的函数签名如下所示:
		func (f *File) WriteAt(b []byte, off int64) (n int, err error)
	以下是对函数签名相关参数说明:
		b:
			待写入的数据内容
		off:
			偏移量(通常是Seek函数返回值)
	*/
	f.WriteAt([]byte("Kafka"), 10)

}
