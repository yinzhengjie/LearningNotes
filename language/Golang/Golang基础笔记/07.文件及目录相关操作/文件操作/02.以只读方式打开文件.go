package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	/**
	打开文件的API函数实现如下所示,很明显它是基于OpenFile实现的
		func Open(name string) (*File, error) {
			return OpenFile(name, O_RDONLY, 0)
		}
	下面是对API的参数解释说明:
		name:
			指的是文件名称,可以是相对路径，也可以是绝对路径。
		*File:
			指的是文件指针,我们使用完文件后要记得释放该文件指针资源
		error:
			指的是创建文件的报错信息，比如指定的文件不存在就会报错"The system cannot find the file specified."
	温馨提示:
		Open()是以只读权限打开文件名为name的文件，得到的文件指针file，只能用来对文件进行"读"操作。
		如果我们有"写"文件的需求，就需要借助"Openfile"函数来打开了。
	*/
	f, err := os.Open("E:\\yinzhengjie\\input\\kafka.txt")

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
		fmt.Println("文件打开成功....")
	}

	/**
	文件成功创建之后,一直处于打开状态，因此我们使用完之后一定要关闭文件，当然关闭文件之前一定要确保为文件已经创建成功。
	*/
	defer f.Close()
}
