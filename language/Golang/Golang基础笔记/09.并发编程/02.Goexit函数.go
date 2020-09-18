package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("Goroutine 666666")

		func() {
			defer fmt.Println("Goroutine 88888888")

			/**
			return、Goexit() 和 os.Exit()的区别:
				return:
					一般用于函数的返回,只能结束当前所在的函数.
				Goexit():
					一般用于协程的退出
					具有击穿特性,能结束掉当前所在的Goroutine，无论存在几层函数调用
				os.Exit():
					主动退出主Goroutine,换句话说，直接终止整个程序的运行。
			*/
			runtime.Goexit() //终止当前Goroutine
			//return
			//os.Exit(100)
			fmt.Println("AAAA")
		}()

		fmt.Println("CCCCC")
	}()

	//我们的主Goroutine会运行15s,有充足时间使得上面的子Go程代码执行完毕哟~
	for index := 1; index <= 30; index += 2 {
		fmt.Printf("Main Say: index = %d\n", index)
		time.Sleep(1 * time.Second)
	}
}
