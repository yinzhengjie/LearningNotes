package main

import (
	"fmt"
)

func main() {

	/*
		switch语法格式如下:

			switch 变量或者表达式的值{
				case 值1:
					//代码块1
				case 值2:
					//代码块2
					//fallthrough
				case 值3:
					//代码块3
				default:
					//代码块4
				}

		执行流程：
			程序执行到switch处，首先将变量或者表达式的值计算出来，然后拿着这个值依次跟每个case后面所带的值进行匹配，一旦匹配成功，则执行该case所对应的代码，执行完成后，跳出switch-case结构。
			如果，跟每个case所带的值都不匹配。就看当前这个switch-case结构中是否存在default，如果有default，则执行default中的语句，如果没有default，则该switch-case结构什么都不做。

		温馨提示：
			(1)某个case后面跟着的代码执行完毕后，不会再执行后面的case，而是跳出整个switch结构， 相当于每个case后面都跟着break(终止)。
			(2)如果想执行完成某个case后，强制执行后面的case，可以使用fallthrough;
			(3)尽量减少float类型作为case分支,因为浮点数精度问题在switch语句中可能会体现出来,比如"3.14"和"3.140000000001"在swich语句中可能会认为是同一个值
	*/

	var (
		month int
		days  int
	)

	fmt.Print("请输入您想要查询的月份对应的天数:>>> ")
	fmt.Scan(&month)

	switch month {
	//如果有多个值都执行同一块代码体,可以将多个值以逗号分割
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		days = 28
	default:
		fmt.Println("请您输入有效的月份[1-12]")
	}

	fmt.Printf("您输入的月份是[%d],该月有[%d]天\n", month, days)
}
