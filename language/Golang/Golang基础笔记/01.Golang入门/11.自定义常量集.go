package main

import (
	"fmt"
)

func main() {
	const (
		RED    = 1
		ORANGE = 2
		YELLOW = 4
		GREEN  = 8
		CYAN   = 16
		BLUE   = 32
		VIOLET = 64
	)

	fmt.Println("赤橙黄绿青蓝紫对应的标记为:", RED, ORANGE, YELLOW, GREEN, CYAN, BLUE, VIOLET)
}
