package main

import (
	"fmt"
	"time"
)

func FibonacciSeriesWrite(fib chan int) {
	a, b := 1, 1
	fmt.Printf("%d\n%d\n", a, b)
	for {
		select {
		case fib <- a + b: //写入数据
			a, b = b, a+b
		}
	}
}

func FibonacciSeriesRead(fib chan int) {
	for {
		fmt.Println(<-fib) //读取数据
		time.Sleep(time.Second)
	}
}

func main() {
	//初始化channel
	s1 := make(chan int)

	go FibonacciSeriesWrite(s1)

	go FibonacciSeriesRead(s1)

	for {
		time.Sleep(time.Second)
	}
}
