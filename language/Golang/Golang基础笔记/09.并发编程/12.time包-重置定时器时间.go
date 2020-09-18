package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(100 * time.Second)

	go func() {
		fmt.Println("当前时间为:", time.Now())

		t := <-timer.C

		fmt.Println("当前时间为:", t)
	}()

	//重置定时器时间
	timer.Reset(3 * time.Second)

	for {
		time.Sleep(1 * time.Second)
	}

}
