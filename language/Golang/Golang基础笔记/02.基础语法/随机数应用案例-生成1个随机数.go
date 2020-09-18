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
		生成一个随机数
	*/
	a := rand.Intn(100) //实际随机生成的数字范围[0,99]
	fmt.Printf("a的类型为[%T],a的随机数值为:[%d]\n", a, a)
}
