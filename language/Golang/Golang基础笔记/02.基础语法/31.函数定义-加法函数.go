package main

import (
	"fmt"
)

func add(a int, b int) (sum int) {
	sum = a + b
	return sum
}

func main() {

	x := 100
	y := 200

	res := add(x, y)

	fmt.Printf("x + y = %d\n", res)
}
