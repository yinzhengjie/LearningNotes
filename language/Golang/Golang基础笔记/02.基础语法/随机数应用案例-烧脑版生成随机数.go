package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomNumber(p *int, start_pos int, stop_pos int) {
	/*
		rand.Seed:
			还函数是用来创建随机数的种子,如果不执行该步骤创建的随机数是一样的，因为默认Go会使用一个固定常量值来作为随机种子。

		time.Now().UnixNano():
			当前操作系统时间的毫秒值
	*/
	rand.Seed(time.Now().UnixNano())

	random_number := rand.Intn(stop_pos-start_pos) + start_pos

	/*
		将指针作为参数传递，形参间接修改实参的值
	*/
	*p = random_number
}

func main() {

	/*
		定义要生成随机值的范围
	*/
	var (
		start_pos int = 200
		stop_pos  int = 500
		num       int
	)

	for index := 1; index <= 10; index++ {
		/*
			每次循环时让CPU休息1毫秒，因为CPU执行太快，可能10此循环几乎在同一纳秒执行完毕。

			我暂时想到的解决方案有两种:
				(1)以形参的方式传入随机种子值
				(2)让主线程在不同的时间段调用函数

			上面两种思路均可，我为了省事情就直接让主线程休眠1毫秒，即采用了第二种解决方案。
		*/
		time.Sleep(1)

		getRandomNumber(&num, start_pos, stop_pos)
		fmt.Printf("第[%d]个随机数的值为:[%d]\n", index, num)
	}

}
