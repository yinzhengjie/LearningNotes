package main

import (
	"fmt"
)

func main() {

	//外层循环，控制循环次数
	for i := 1; i <= 9; i++ {
		//内层循环，控制每次外层循环，内层循环执行的次数
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
