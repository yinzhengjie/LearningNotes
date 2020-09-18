package main

import (
	"fmt"
)

func main() {

	var score int
	fmt.Printf("请问你的分数是多少:>>> ")
	fmt.Scan(&score)

	/*
		if嵌套结构语法格式如下:
			if 条件判断 {
				 if 条件判断 {
					//代码块1
				}
				...
			}else if 条件判断 {
				if 条件判断 {
					//代码块2
				}
				...
			}
			...
			}else{
				//代码块2
			}

		if多分支代码块必须且只有一个代码块会被执行:
			从上到下依次判断条件，如果结果为真，就执行符合相应条件判断内的代码块语句。
	*/

	if score >= 710 {
		fmt.Println("我要买电脑")
		if score >= 735 {
			fmt.Println("我要买钢琴")
		}
	} else if score >= 650 {
		fmt.Println("我要买手机")
		if score >= 680 {
			fmt.Println("我要学习计算机技术")
		}
	} else {
		fmt.Println("我要吃冰棍")
	}
}
