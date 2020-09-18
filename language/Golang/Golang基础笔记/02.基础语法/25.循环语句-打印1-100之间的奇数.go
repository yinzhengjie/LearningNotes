package main

import (
	"fmt"
)

func main() {

	//计算1-100之间的奇数之和
	var sum int

	for i := 1; i <= 100; i += 2 {
		sum += i
	}

	fmt.Printf("[1-100]之间数字之和为:%d", sum)

}
