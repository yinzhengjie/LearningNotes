package main

import (
	"fmt"
)

func f1() (int, error) {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return fmt.Println(4)
}

func f2() (int, error) {
	defer fmt.Println(5)
	defer fmt.Println(6)
	fmt.Println(7)
	f1()

	defer func() (int, error) {
		defer fmt.Println(8)
		fmt.Println(9)
		return fmt.Println(10)
	}()

	return fmt.Println(11)
}

func main() {
	f2()
}
