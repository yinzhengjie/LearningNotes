package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	/**
	创建文件的API函数签名如下所示:
		func Create(name string) (*File, error)
	下面是对API的参数解释说明:
		name:
			指的是文件名称,可以是相对路径，也可以是绝对路径。
		*File:
			指的是文件指针,我们使用完文件后要记得释放该文件指针资源
		error:
			指的是创建文件的报错信息，比如指定的文件父目录不存在就会报错"The system cannot find the path specified."
	温馨提示:
		根据提供的文件名创建新的文件，返回一个文件对象，返回的文件对象是可读写的。
		创建文件时，如果存在重名的文件就会覆盖掉原来的文件。换句话说,如果文件存在就清空文件内容并打开新文件,如果文件不存在则创建新文件并打开。
	*/
	f, err := os.Create("E:\\yinzhengjie\\input\\kafka.txt")

	/**
	一旦文件报错就执行return语句，下面的defer语句就不会被执行到哟~

	温馨提示:
		我们不要将下面的defer语句和判断错误的语句互换位置，因为判断错误的语句是来确保文件是否创建成功;
		如果有错误意味着文件没有被成功创建,换句话说,如果文件创建失败，那么文件指针为空，此时如果执行关闭文件会报错哟;
		如果没有错误就意味着文件创建成功，即在执行关闭文件的操作是确保不会报错
	*/
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
}
