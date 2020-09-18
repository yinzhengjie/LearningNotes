package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("当前时间为:", time.Now())

	//阻塞2s后产生一个事件，往channel写内容
	<-time.After(3 * time.Second)

	fmt.Println("当前时间为:", time.Now())

	for {
		time.Sleep(1 * time.Second)
	}
}
