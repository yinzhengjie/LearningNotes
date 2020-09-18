package main

import (
	"fmt"
	"strconv"
	"time"
)

//定义生产者，假设生产者不消费
func Producer(p chan<- string) {
	defer fmt.Println("生产蛋糕结束")
	for index := 1; index <= 10; index++ {
		p <- "生产了" + strconv.Itoa(index) + "个蛋糕."
	}
}

//定义消费者,假设消费者不生产
func Consumer(c <-chan string) {
	defer fmt.Println("吃饱了")
	for index := 1; index <= 8; index++ {
		fmt.Println(<-c)
	}
}

func main() {
	s1 := make(chan string, 10)

	go Producer(s1)

	go Consumer(s1)

	for {
		time.Sleep(1 * time.Second)
	}
}
