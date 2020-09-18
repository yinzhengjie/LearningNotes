package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex //定义互斥锁

func MyPrint(data string) {
	//mutex.Lock()         //添加互斥锁
	//defer mutex.Unlock() //使用结束时自动解锁

	for _, value := range data { //迭代字符串的每个字符并打印
		fmt.Printf("%c", value)
		time.Sleep(time.Second) //模拟Go程在执行任务
	}
	fmt.Println()
}

func Show01(s1 string) {
	MyPrint(s1)
}

func Show02() {
	MyPrint("Jason Yin")
}

func main() {
	/**
	虽然我们在主Go中开启了2个子Go程，但由于2个子Go程有互斥锁的存在,因此一次只能运行一个Go程哟~
	*/
	go Show01("尹正杰")
	go Show02()
	time.Sleep(time.Second * 30) //主Go程设置充足的时间让所有子Go程执行完毕~因为主Go程结束会将所有的子Go程杀死。
}
