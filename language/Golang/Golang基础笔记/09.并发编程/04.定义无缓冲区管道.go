package main

import (
	"fmt"
	"time"
)

func Read(s chan int) {
	defer fmt.Println("读端结束~~~")
	for index := 0; index < 3; index++ {
		fmt.Printf("读取到channel的数据是: %d\n", <-s)
	}
}

func Write(s chan int, value int) {
	defer fmt.Println("写端结束...")
	for index := 100; index < value; index++ {
		s <- index
	}
}

func main() {

	/**
	无缓冲区channel特性:
		只有在写段和读端同时准备就绪的情况下才能运行。
	*/
	s1 := make(chan int) //定义一个无缓冲区的channel
	//s1 := make(chan int, 10) //定义一个有缓冲区的channel,其容量为10

	go Read(s1)

	go Write(s1, 110)

	/**
	我们让主Goroutine阻塞，写个死循环即可.
	*/
	for {
		time.Sleep(1 * time.Second)
	}
}
