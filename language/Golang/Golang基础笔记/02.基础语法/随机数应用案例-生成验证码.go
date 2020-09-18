package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/*
		rand.Seed:
			还函数是用来创建随机数的种子,如果不执行该步骤创建的随机数是一样的，因为默认Go会使用一个固定常量值来作为随机种子。

		time.Now().UnixNano():
			当前操作系统时间的毫秒值
	*/
	rand.Seed(time.Now().UnixNano())

	/*
		生成十个随机数[200-500]
	*/
	for index := 1; index <= 10; index++ {
		random_number := rand.Intn(301) + 200
		fmt.Printf("第[%d]个随机数的值为:[%d]\n", index, random_number)
	}
}
