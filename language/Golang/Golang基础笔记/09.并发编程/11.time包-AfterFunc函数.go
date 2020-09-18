package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("当前时间为:", time.Now())

	//延迟5s后调用函数
	time.AfterFunc(time.Second*5, func() {
		fmt.Println("当前时间为:", time.Now())
		fmt.Println("In AfterFunc ...")
	})

	for {
		time.Sleep(1 * time.Second)
	}
}
