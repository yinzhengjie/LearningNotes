package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(5 * time.Second)

	go func() {
		fmt.Println("当前时间为:", time.Now())

		t := <-timer.C

		fmt.Println("当前时间为:", t)
	}()

	//停止计时器
	timer.Stop()

	for {
		time.Sleep(1 * time.Second)
	}

}
