package main

import (
	"fmt"
)

func main() {

	/*
		单分支结构语法格式如下:
			if 条件判断 {
				//代码块
			}

		条件判断如果为真（true），那么就执行大括号中的代码块；如果为假（false），就不执行大括号中的代码块。
	*/

	var score int

	fmt.Printf("请输入你的分数:>>> ")
	fmt.Scan(&score)

	if score >= 730 {
		fmt.Println("恭喜你，你有上哈佛的潜质...")
	}
}
