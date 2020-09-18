package main

import (
	"fmt"
	"time"
)

func Task(start int, end int, desc string) {
	for index := start; index <= end; index += 2 {
		fmt.Printf("%s %d\n", desc, index)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	/**
	创建Goroutine:
		只需在函数调⽤语句前添加Go关键字，就可创建并发执⾏单元。开发⼈员无需了解任何执⾏细节，调度器会自动将其安排到合适的系统线程上执行。
		在并发编程中，我们通常想将一个过程切分成几块，然后让每个goroutine各自负责一块工作，当一个程序启动时，主函数在一个单独的goroutine中运行，我们叫它main goroutine。新的goroutine会用go语句来创建。而go语言的并发设计，让我们很轻松就可以达成这一目的。

	Goroutine特性:
		为了避免类似孤儿进程的存在，如果main协程挂掉，所有协程都挂掉。
		换句话说,主goroutine退出后，其它的工作goroutine也会自动退出。
	*/
	go Task(10, 30, "Task Func Say: index =")

	Task(11, 30, "Main Say: index =")
}
