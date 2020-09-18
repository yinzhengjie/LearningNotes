package main

import (
	"fmt"
)

func main() {
	PI := 3.14
	r := 4.115
	fmt.Printf("变量PI的数据类型是:%T\n", PI)
	fmt.Printf("变量r的数据类型是:%T\n", r)
	//圆的的面积
	s := PI * r * r

	//圆的周长
	l := PI * 2 * r

	/*
		带小数的数据类型称为浮点数，浮点数分为单精度浮点型(float32)和双精度浮点型(float64):
			单精度浮点型(float32):
				小数点默认保留6位
			双精度浮点型(float64)
				小数点默认保留15位

		关于浮点数的知识扫描，博主推荐阅读:
			https://www.cnblogs.com/yinzhengjie2020/p/12247502.html
	*/
	fmt.Println("圆的面积:", s)
	fmt.Println("圆的周长:", l)

	fmt.Printf("%.2f\n", l) //保留2位有效数字
	fmt.Printf("%.5f\n", l) //保留5位有效数字
}
