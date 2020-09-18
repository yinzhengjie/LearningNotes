package main //命令行源码文件必须在这里声明自己属于main包

/*
   使用import关键字导入包，建议每导入一个包占用一行，看起来比较美观。
*/
import (
	"fmt"
)

func main() {
	fmt.Println("hello world") //打印字符串并换行
}
