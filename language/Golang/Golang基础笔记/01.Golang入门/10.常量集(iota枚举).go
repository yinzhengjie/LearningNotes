package main

import (
	"fmt"
)

func main() {

	//在同一个常量集中，第一个iota等于0，每当iota在新的一行被使用时，它的值会自动加1
	const (
		A = iota
		B = iota
		C = iota
	)

	fmt.Printf("A = %d\nB = %d\nC = %d\n\n", A, B, C)

	//在同一个常量集中，第一个iota等于0，每当iota在新的一行被使用时，它的值会自动加1,下面这种写法是上面写法的简写形式
	const (
		A1 = iota
		B1
		C1
	)

	fmt.Printf("A1 = %d\nB1 = %d\nC1 = %d\n\n", A1, B1, C1)

	//在同一个常量集中，第一个iota等于0，每当iota在新的一行被使用时，它的值会自动加1,但是同一行的iota值是相同的哟~
	const (
		A2     = iota
		B2, C2 = iota, iota
		D2     = iota
	)

	fmt.Printf("A2 = %d\nB2 = %d\nC2 = %d\nD2 = %d\n\n", A2, B2, C2, D2)

}
