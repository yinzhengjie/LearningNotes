package main

import (
	"fmt"
)

func main() {

	/*
		标识符的命名规则如下:
			(1)允许使用字母、数字、下划线
		　　(2)不允许使用Go语言关键字
		　　(3)不允许使用数字开头
		　　(4)区分大小写

		满足上面的Go编译器的要求后，生产环境中推荐命名规则:
		　　(1)见名知义
		　　(2)驼峰命名法
				小驼峰式命名法（lower camel case）：
		　　　　		第一个单词以小写字母开始，从第二个单词开始首字母大写，例如：myNginxPort
		　　　　大驼峰式命名法（upper camel case）：
		　　　　　　每一个单字的首字母都采用大写字母，例如：FirstName、LastName
		　　(3)下划线命名法
		　　　　　	每个单次都小写，各单次之间使用下划线进行分割，例如:my_cluster
	*/

	//小驼峰命名
	myNginxPort := "node101.yinzhengjie.org.cn:80"
	fmt.Println(myNginxPort)

	//大驼峰命名
	FirstName := "yin"
	LastName := "zhengjie"
	fmt.Println(FirstName)
	fmt.Println(LastName)

	//下划线命名
	my_cluster := "yinzhengjie_bigdata"
	fmt.Println(my_cluster)

}
