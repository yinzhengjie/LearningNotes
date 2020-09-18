package main

import (
	"fmt"
)

func main() {

	var arr []rune = []rune{'湾', '前', '过', '渡', '小', '舟', '虚'}

	for index := 0; index <= len(arr)-1; index++ {
		fmt.Print(string(arr[index]))
	}

	fmt.Println()

	for index := len(arr) - 1; index >= 0; index-- {
		fmt.Print(string(arr[index]))
	}
}
