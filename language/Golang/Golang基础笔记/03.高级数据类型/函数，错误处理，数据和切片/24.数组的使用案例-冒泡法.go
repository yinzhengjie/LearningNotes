package main

import (
	"fmt"
)

func main() {

	arr := [10]int{1, 9, 8, 5, 6, 7, 0, 4, 3, 2}
	fmt.Printf("arr数组未经过排序之前,元素顺序为:%d\n", arr)

	//外层循环控制行，表示执行次数
	for i := 0; i < len(arr)-1; i++ {
		//内层循环控制列，表示比较次数
		for j := 0; j < len(arr)-1-i; j++ {
			//比较两个相邻元素
			if arr[j] > arr[j+1] {
				//交换数据
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Printf("arr数组经过排序之后,元素顺序为:%d\n", arr)
}
