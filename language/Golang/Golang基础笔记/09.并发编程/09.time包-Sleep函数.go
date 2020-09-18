package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("当前时间为:", time.Now())

	time.Sleep(3 * time.Second)

	fmt.Println("当前时间为:", time.Now())

}
