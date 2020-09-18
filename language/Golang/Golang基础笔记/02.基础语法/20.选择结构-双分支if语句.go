package main

import (
	"fmt"
)

func main() {

	var score int
	fmt.Printf("请问你的分数是多少:>>> ")
	fmt.Scan(&score)

	/*
		单分支结构语法格式如下:
			if 条件判断 {
				//代码块1
			}else{
				//代码块2
			}

		if代码块或else代码块，必须且只有一个代码块会被执行:
			条件判断如果为真（true），那么就执行if大括号中的语句;
			条件判断如果为假（false），那么就执行else大括号中的语句;
	*/
	if score >= 680 {
		fmt.Println("恭喜你，你有上清华的潜质...")
	} else {
		fmt.Println("其实你可以先考虑其它学校~")
	}
}
