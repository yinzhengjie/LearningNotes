package main

import "fmt"

func main() {

	var (
		sum int
	)

	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
	}

	fmt.Printf("1-100之间的偶数之和为:%d\n", sum)
}
