package main

import (
	"fmt"
	"time"
)

func main() {

	//注意，无缓冲区channel在读端和写段都准备就绪的时候不阻塞
	s1 := make(chan int)

	/**
	子线程读取:
		先开启一个子Go程用于读取无缓冲channel中的数据,此时由于写端未就绪因此子Go程会处于阻塞状态,但并不会影响主Go程,因此代码可以继续向下执行哟~
	*/
	go func() {
		fmt.Println(<-s1)
	}()

	/**
	主线程写入:
		此时读端(子Go程)处于阻塞状态正在准备读取数据,主Go程在写入数据时,子Go程会立即消费掉哟~
	*/
	s1 <- 5

	for {
		time.Sleep(time.Second)
	}
}
