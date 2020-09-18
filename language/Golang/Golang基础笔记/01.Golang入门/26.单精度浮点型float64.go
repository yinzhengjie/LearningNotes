package main

import (
	"fmt"
)

func main() {
	/*
		温馨提示:
			自动推导类型命名的浮点数类型默认为float64.
	*/
	a := 3.141592653589
	var b float64 = 1.618033988749
	/*
		3.141592653589 x 1.618033988749 = 5.083203592311165

		但是单精度浮点型小数点后默认保留六位，最终结果四舍五入法得到的结果为:5.083204
	*/
	var c float64 = a * b

	fmt.Printf("a的类型为:[%T],a的值为:[%f]\n", a, a)
	fmt.Printf("b的类型为:[%T],b的值为:[%f]\n", b, b)
	fmt.Printf("c的类型为:[%T],c的值为:[%f]\n", c, c)

	fmt.Printf("a x b = %f\n", c)

	/*
		%.1f:
			表示小数点后按照四舍五入法保留1位有效数字
		%.2f:
			表示小数点后按照四舍五入法保留2位有效数字
		%.3f:
			表示小数点后按照四舍五入法保留3位有效数字
		...:
			综上所述，依此类推即可....
	*/
	fmt.Printf("a x b = %.1f\n", c)
	fmt.Printf("a x b = %.2f\n", c)
	fmt.Printf("a x b = %.3f\n", c)
}
