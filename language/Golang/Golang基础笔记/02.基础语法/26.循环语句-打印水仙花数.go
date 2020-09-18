package main

import (
	"fmt"
	"math"
)

func main() {

	/*
		什么是水仙花数:
			一个三位数，各个位数的立方和等于本身的数统称为水仙花数字

		关于math包的用法，博主推荐阅读:
			https://www.cnblogs.com/yinzhengjie/p/12203765.html
	*/

	fmt.Printf("1000以内的水仙花数有: ")
	//计算1000以内的水仙花数
	for i := 100; i <= 999; i++ {
		//计算个位的三次方
		one := int(math.Pow(float64(i%10), 3))

		//计算十位的三次方
		ten := int(math.Pow(float64(i/10%10), 3))

		//计算百位的三次方
		hundred := int(math.Pow(float64(i/100), 3))

		if one+ten+hundred == i {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

}
