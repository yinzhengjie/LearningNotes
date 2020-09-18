package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	s1 := make(chan int, 1)

	go func() {
		for {
			select {
			case s1 <- 110:
				fmt.Println("写入channel数据")
				fmt.Println("当前时间为:", time.Now())
			/**
			设置定时器，当channel中的数据在30秒内没有被消费时,就会退出程序。
			*/
			case <-time.After(time.Second * 30):
				fmt.Println("程序响应超过30秒,程序已退出!")
				fmt.Println("当前时间为:", time.Now())
				os.Exit(100)
			}
		}
	}()

	/**
	设置一次性定时器,仅消费一次数据哟~
	*/
	time.AfterFunc(
		time.Second*3,
		func() {
			fmt.Printf("获取到channel中的数据为: %d\n", <-s1)
		})

	for {
		time.Sleep(time.Second)
	}
}
