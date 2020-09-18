package main

import (
	"fmt"
	"runtime"
	"time"
)

func ReadChannel(s chan int) {
	defer fmt.Println("读端结束~~~")
	for index := 0; index < 5; index++ {
		if index == 103 {
			/**
			程序结束的时候，清理掉channel占用的空间，只影响写入，不影响读取。
			*/
			close(s)
		}

		/**
		从channel中接收数据，并赋值给value,同时检查通道是否已关闭或者是否为空
		*/
		value, ok := <-s
		if !ok { //等效于"ok != true"
			fmt.Println("channel已关闭或者管道中没有数据")
			runtime.Goexit()
		} else {
			fmt.Printf("读取到channel的数据是: %d\n", value)
		}
	}
}

func WriteChannel(s chan int, value int) {
	defer fmt.Println("写端结束...")
	for index := 100; index < value; index++ {
		if index == 103 {
			/**
			程序结束的时候，清理掉channel占用的空间，只影响写入，不影响读取。
			*/
			close(s)
		}
		s <- index
	}
}

func main() {

	s1 := make(chan int) //定义一个有缓冲区的channel,其容量为10

	go ReadChannel(s1)

	go WriteChannel(s1, 110)

	for {
		time.Sleep(1 * time.Second)
	}
}
