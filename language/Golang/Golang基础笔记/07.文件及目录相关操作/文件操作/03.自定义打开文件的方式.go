package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	/**
	  创建文件的API函数签名如下所示:
	      func OpenFile(name string, flag int, perm FileMode) (*File, error)

	  下面是对API的参数解释说明:
	      name:
	          指的是文件名称,可以是相对路径，也可以是绝对路径。
	      flag:
	          表示读写模式，常见的模式有:O_RDONLY(只读模式), O_WRONLY(只写模式), O_RDWR(可读可写模式)。
	      perm:
	          表示打开权限。来源于Linux系统调用中的open函数，参2为：O_CREATE时，可创建新文件。
	          权限取值范围是八进制,即"0-7".表示如下：
	              0：没有任何权限
	              1：执行权限(如果是可执行文件，是可以运行的)
	              2：写权限
	              3: 写权限与执行权限
	              4：读权限
	              5: 读权限与执行权限
	              6: 读权限与写权限
	              7: 读权限，写权限，执行权限
	      *File:
	          指的是文件指针,我们使用完文件后要记得释放该文件指针资源
	      error:
	          指的是创建文件的报错信息，比如指定的文件父目录不存在就会报错"The system cannot find the path specified."

		温馨提示:
			使用OpenFile打开的文件，默认写的时候是聪文件开头开始写入数据,这样会将原来的数据覆盖掉;
			如果不想覆盖的方式写入,使用追加的方式写入,具体代码如下所示(当然你也可以使用Seek函数实现哟~)。
	*/
	f, err := os.OpenFile("E:\\yinzhengjie\\input\\kafka.txt", os.O_RDWR|os.O_APPEND, 0666)

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
