package main

import (
	"fmt"
)

func BubbleSort(s1 []int) {
	for i := 0; i < len(s1)-1; i++ {
		for j := 0; j < len(s1)-1-i; j++ {
			if s1[j] > s1[j+1] {
				s1[j], s1[j+1] = s1[j+1], s1[j]
			}
		}
	}
}

func main() {

	arr := []int{1, 9, 8, 5, 6, 7, 0, 4, 3, 2}
	fmt.Printf("arr数组未经过排序之前,元素顺序为:%d\n", arr)

	BubbleSort(arr)

	fmt.Printf("arr数组未经过排序之前,元素顺序为:%d\n", arr)

}
