package main

import (
	"fmt"
	"time"
)

func main() {

	//设置定时器为3秒
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("当前时间为:", time.Now())

	t := <-timer.C //从定时器拿数据
	fmt.Println("当前时间为:", t)

}
