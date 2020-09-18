package main

import (
	"fmt"
	"runtime"
)
import "sync"
import "math/rand"
import "time"

/**
创建全局条件变量
*/
var cond sync.Cond

// 生产者
func producer(out chan<- int, idx int) {
	for {
		/**
		条件变量对应互斥锁加锁,即在生产数据时得加锁。
		*/
		cond.L.Lock()

		/**
		产品区满3个就等待消费者消费
		*/
		for len(out) == 3 {
			/**
			挂起当前go程， 等待条件变量满足，被消费者唤醒,该函数的作用可归纳为如下三点：
				1>.阻塞等待条件变量满足
				2>.释放已掌握的互斥锁相当于cond.L.Unlock()。注意：两步为一个原子操作。
				3>.当被唤醒，Wait()函数返回时，解除阻塞并重新获取互斥锁。相当于cond.L.Lock()
			*/
			cond.Wait()
		}

		/**
		产生一个随机数,写入到 channel中(模拟生产者）
		*/
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("%dth 生产者，产生数据 %3d, 公共区剩余%d个数据\n", idx, num, len(out))

		/**
		单发通知，给一个正等待(阻塞)在该条件变量上的goroutine(Go程)发送通知。换句话说,唤醒阻塞的消费者
		*/
		//cond.Signal()

		/**
		广播通知，给正在等待（阻塞）在该条件变量上的所有goroutine(线程)发送通知。
		*/
		cond.Broadcast()

		/**
		生产结束，解锁互斥锁
		*/
		cond.L.Unlock()

		/**
		生产完休息一会，给其他Go程执行机会.
		*/
		time.Sleep(time.Second)
	}
}

//消费者
func consumer(in <-chan int, idx int) {
	for {
		/**
		条件变量对应互斥锁加锁(与生产者是同一个)
		*/
		cond.L.Lock()

		/**
		产品区为空 等待生产者生产
		*/
		for len(in) == 0 {
			/**
			挂起当前go程， 等待条件变量满足，被生产者唤醒
			*/
			cond.Wait()
		}

		/**
		将channel中的数据读走(模拟消费数据)
		*/
		num := <-in
		fmt.Printf("[%dth] 消费者, 消费数据 %3d,公共区剩余%d个数据\n", idx, num, len(in))
		/**
		唤醒阻塞的生产者
		*/
		cond.Signal()
		/**
		消费结束，解锁互斥锁
		*/
		cond.L.Unlock()

		/**
		消费完休息一会，给其他Go程执行机会
		*/
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	/**
	设置随机数种子
	*/
	rand.Seed(time.Now().UnixNano())

	/**
	产品区(公共区)使用channel模拟
	*/
	product := make(chan int, 3)

	/**
	创建互斥锁和条件变量(申请内存空间)
	*/
	cond.L = new(sync.Mutex)

	/**
	创建3个生产者
	*/
	for i := 101; i < 103; i++ {
		go producer(product, i)
	}

	/**
	创建5个消费者
	*/
	for i := 211; i < 215; i++ {
		go consumer(product, i)
	}
	for { // 主go程阻塞 不结束
		runtime.GC()
	}
}
