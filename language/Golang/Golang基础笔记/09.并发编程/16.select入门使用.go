package main

import (
	"fmt"
	"time"
)

func main() {

	s1 := make(chan int, 1)
	//s1 := make(chan int, 0)	//无缓冲的channel

	number := 1

	for {
		/**
		使用select关键字来监听指定channel的读写情况
		*/
		select {
		case s1 <- number:
			fmt.Println("奇数: ", number)
			number++
			time.Sleep(time.Second * 1)

		case <-s1:
			fmt.Println("偶数: ", number)
			number++
			time.Sleep(time.Second * 1)
			/**
			当读取和写入(即I/O操作)都不满足的情况下，就会执行默认的条件，需要将channel的容量设置为0就可以看到效率。
			*/
		default:
			fmt.Println("======")
			time.Sleep(time.Second * 1)
		}
	}
}
