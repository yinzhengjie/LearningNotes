package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 6; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
}
