package main

import "fmt"

func main() {

	var index int

	//使用for的死循环语法
	for {
		index++
		fmt.Printf("%d\t", index)

		//打印数字1-10
		if index == 10 {
			break
		}
	}
	fmt.Println()
}
