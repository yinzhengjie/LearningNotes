package main

import (
	"fmt"
	"os"
	"syscall"
)

const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
)

func main() {
	f, err := os.OpenFile("E:\\yinzhengjie\\input\\flume.txt", O_RDWR|O_CREATE|O_TRUNC, 0666)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()

	f.Write([]byte("Flume 是一个文本日志收集工具。\n"))
	f.WriteString("Flume是一种分布式，可靠且可用的服务，用于有效地收集，聚合和移动大量日志数据。")
	f.WriteAt([]byte("Flume具有基于流数据流的简单灵活的体系结构。"), 60) //可以指定偏移量写入

	/**
	Seek的函数签名如下所示:
		func (f *File) Seek(offset int64, whence int) (ret int64, err error)
	以下是对函数签名相关参数说明:
		offset:
			指定偏移量,如果是正数,则表示向文件尾偏;如果是负数,则表示向文件头偏.
		whence:
			指定偏移量的起始位置.
			io.SeekStart: 文件起始位置,对应常量整型0.
			io.SeekCurrent： 文件当前位置,对应常量整型1.
			io.SeekEnd: 文件结尾位置,对应常量整型2.
		返回值:
			表示从文件起始位置，到当前文件读写指针位置的偏移量。
	*/
	f.Seek(10, 2)

	f.WriteString("Flume具有可调整的可靠性机制以及许多故障转移和恢复机制，具有强大的功能和容错能力。")
}
