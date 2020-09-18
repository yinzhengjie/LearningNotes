package main

import (
	"fmt"
)

func main() {

	var score int
	fmt.Printf("请问你的语文分数是多少:>>> ")
	fmt.Scan(&score)

	/*
		单分支结构语法格式如下:
			if 条件判断1 {
				//代码块1
			}else if 条件判断2{
				//代码块2
			}else if 条件判断3{
				//代码块3
			...
			}
			}else{
				//代码块4
			}

		if多分支代码块必须且只有一个代码块会被执行:
			从上到下依次判断条件，如果结果为真，就执行符合相应条件判断内的代码块语句。
	*/

	if score > 150 {
		fmt.Println("请输入合法的分数[0-150]")
	} else if score >= 120 {
		fmt.Println("你是尖子生,你的稳住成绩别下滑")
	} else if score >= 90 {
		fmt.Println("你的成绩还有很大的提升空间,加油~")
	} else {
		fmt.Println("小伙子你得努力啦,不然家长会不好过啊~")
	}
}
