package main

import (
	"fmt"
	"time"
)

/**
s的类型说明:
	"chan<- int"表示传入只写的管道
*/
func Send(s chan<- int, value int) {
	s <- value
}

/**
r的类型说明:
	"<-chan int"表示传入只读的管道
*/
func Receive(r <-chan int) {
	fmt.Printf("管道中的数据为:%d\n", <-r)
}

func main() {

	//创建管道
	s1 := make(chan int, 5)

	go Receive(s1)

	go Send(s1, 110)

	for {
		time.Sleep(1 * time.Second)
	}
}
