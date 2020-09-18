package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	number int
	rwlock sync.RWMutex //定义读写锁
)

func MyRead(n int) {
	//rwlock.RLock()         //添加读锁
	//defer rwlock.RUnlock() //使用结束时自动解锁
	fmt.Printf("[%d] Goroutine读取数据为: %d\n", n, number)
}

func MyWrite(n int) {
	rwlock.Lock()         //添加写锁
	defer rwlock.Unlock() //使用结束时自动解锁
	number = rand.Intn(100)
	fmt.Printf("%d Goroutine写入数据为: %d\n", n, number)
}

func main() {

	//创建写端
	for index := 201; index <= 205; index++ {
		go MyWrite(index)
	}

	//创建读端
	for index := 110; index <= 130; index++ {
		go MyRead(index)
	}

	for {
		time.Sleep(time.Second)
	}
}
