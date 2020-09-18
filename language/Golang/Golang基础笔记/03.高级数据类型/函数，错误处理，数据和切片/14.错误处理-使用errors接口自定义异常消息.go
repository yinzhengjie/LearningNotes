package main

import (
	"errors"
	"fmt"
)

func div(x int, y int) (res int, err error) {

	if y == 0 {
		/*
			Go语言引入了一个关于错误处理的标准模式，即error接口，该接口需要导入"errors"包:
				errors接口适合返回可控的错误，即我们知道在某个代码块中可能会出现的异常。
		*/
		err = errors.New("div函数除零异常...")
		return
	}
	res = x / y
	return
}

func main() {
	a := 10
	b := 0

	/*
		res:
			用来接收运算结果
		err:
			用来接收错误消息
		nil:
			表示空，"if err != nil"表示变量err接收的错误消息是否为空。
	*/
	res, err := div(a, b)
	if err != nil {
		fmt.Println(err)
		//log.Fatal(err) //我们也可以使用log包来帮助咱们输出错误，它会在错误消息签名自动加上日期时间
	} else {
		fmt.Println(res)
	}

	fmt.Println("代码执行完毕")
}
