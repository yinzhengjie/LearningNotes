package main

import (
	"fmt"
)

func main() {

	/*
		从一个整型数种取出最大的整数，最小的整数，数组个元素的总和及平均值
	*/
	var arr [5]int = [5]int{100, 80, 300, 500, 1024}

	var (
		max   = arr[0]
		min   = arr[0]
		total = 0
	)

	for index := 1; index < len(arr); index++ {
		//取最大值
		if arr[index] > max {
			max = arr[index]
		}
		//取最小值
		if arr[index] < min {
			min = arr[index]
		}
		//求和
		total += arr[index]
	}

	//求平均数
	average := total / len(arr)

	fmt.Printf("数组中最大值为:[%d]\n数组中最小值为:[%d]\n数组中总和为:[%d]\n数组中的平均数为:[%d]\n", max, min, total, average)
}
