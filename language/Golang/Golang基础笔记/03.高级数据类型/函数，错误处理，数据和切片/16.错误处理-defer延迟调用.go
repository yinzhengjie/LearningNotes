package main

import (
	"fmt"
)

func main() {

	/*
		defer语句被用于预定对一个函数的调用。可以把这类被defer语句调用的函数称为函数。

		defer的应用场景:
			(1)释放占用的资源;
			(2)捕获异常处理
			(3)输出日志

		温馨提示:
			如果一个函数中有多个defer语句，它们会以LIFO(后进先出)的顺序执行。

		先不要运行代码，你先猜猜下面的代码的执行顺序。
	*/
	defer fmt.Println("\thttps://www.cnblogs.com/yinzhengjie/")

	fmt.Println("我的名字叫尹正杰,英文名叫'Jason Yin'.爱好开源技术.")

	defer fmt.Println("\thttps://www.cnblogs.com/yinzhengjie2020")

	fmt.Println("我的博客是:")
}
