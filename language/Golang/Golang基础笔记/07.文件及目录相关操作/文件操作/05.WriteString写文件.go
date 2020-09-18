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
	写入字符串到文件，其函数实现如下,很明显,WriteString底层调用的依旧是Write方法。
		func (f *File) WriteString(s string) (n int, err error) {
			return f.Write([]byte(s))
		}
	如上所示，我们只需要传入需要写入的字符串即可将内容写到操作系统中对应的文件。
	*/
	f.WriteString("Kafka是一个消息队列")

}
