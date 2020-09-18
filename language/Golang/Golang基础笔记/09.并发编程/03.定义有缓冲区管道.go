package main

import "fmt"

func main() {
	/**
	和map类似，channel也一个对应make创建的底层数据结构的引用，创建channel语法格式如下所示:
		make(chan Type,capacity)

	以下是相关参数说明:
		chan:
			是创建channel所需使用的关键字。
		Type:
			是指定channel收发数据的类型。
		capacity:
			是指定channel的大小。
			当参数"capacity = 0"时,channel是无缓冲阻塞读写的，即该channel无容量。
			当参数"capacity > 0"时,channel有缓冲,是非阻塞的,直到写满capacity个元素才阻塞写入。
	*/
	s1 := make(chan int, 3) //定义一个有缓冲区的channel

	//向channel写入数据
	s1 <- 110
	s1 <- 119
	s1 <- 120

	/**
	注意哈，由于上面已经写了3个数据了，此时s1这个channel的容量已经达到3个容量上限啦,即该channel已满.
	如果channel已满继续往该channel写入数据则会抛出异常:"fatal error: all goroutines are asleep - deadlock!"
	*/
	//s1 <- 114

	//从channel读取数据
	fmt.Println(<-s1)
	fmt.Println(<-s1, <-s1)

	/**
	上面的代码已经对channel进行读写各三次，此时该channel中并没有数据啦。
	如果channel无数据，我们从该channel读取数据时也会抛出异常:"fatal error: all goroutines are asleep - deadlock!"
	*/
	//fmt.Println(<-s1)
}
