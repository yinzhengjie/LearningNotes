package main

import (
	"fmt"
	"os"
)

func main() {

	/**
	  如下所示，打开目录和打开文件的函数是同一个函数哟:
	      func OpenFile(name string, flag int, perm FileMode) (*File, error)

	  函数签名各个参数解释如下:
	      name:
	          表示要打开的目录名称。使用绝对路径较多
	      flag:
	          表示打开目录的读写模式,通常传:O_RDONLY(只读模式)
	      perm:
	          表示打开权限。但对于目录来说略有不同。通常传os.ModeDir。
	      返回值:
	          由于是操作目录，所以file是指向目录的文件指针(*File)。
	          error中保存错误信息。
	*/
	f, err := os.OpenFile("E:\\yinzhengjie\\input", os.O_RDONLY, os.ModeDir)

	if err != nil {
		fmt.Println("目录打开失败: ", err)
		return
	}
	defer f.Close()

	/**
	    Readdir函数的函数签名如下:
	        func (f *File) Readdir(n int) ([]FileInfo, error)

	    函数签名各个参数解释如下:
	        n:
	            Readdir读取与文件关联的目录的内容，并按目录顺序返回由Lstat返回的最多n个FileInfo值组成的片段。对同一文件的后续调用将产生更多的文件信息。
	            如果n>0，Readdir最多返回n个FileInfo结构。在这种情况下，如果Readdir返回一个空片段，它将返回一个非nil错误来解释原因。在目录的末尾，错误是io.EOF。
	            如果n<=0，Readdir将在一个切片中返回目录中的所有文件信息。在这种情况下，如果Readdir成功（一直读取到目录的末尾），则返回切片和nil错误。如果在目录结束之前遇到错误，Readdir将返回在该点之前读取的FileInfo和非nil错误。
	        返回值:
	            返回两个值，一个是读取到的文件信息切片对象("[]FileInfo"),另一个是错误信息("error")

			温馨提示:
				FileInfo中可以获取到问问你家的名称，大小，权限，修改时间,是否是目录等。
	*/
	files, err := f.Readdir(-100)
	if err != nil {
		fmt.Println("错误信息: ", err)
		return
	}

	for _, file := range files {
		fmt.Printf("文件名称是: %s,文件大小是:%d,是否是目录:%t\n", file.Name(), file.Size(), file.IsDir())
	}
}
