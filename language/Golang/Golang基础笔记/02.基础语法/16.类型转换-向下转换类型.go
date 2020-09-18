package main

import (
	"fmt"
)

func main() {

	/*
		疫情期间我家周围的物美超市大白菜的价格竟然涨到了2.75元人民币/500g，你们呢？
	*/
	Chinese_cabbage_price := 2.75 //自动推导浮点型默认是float64类型
	weight := 5                   //自动推导整形默认是int类型
	fmt.Printf("Chinese_cabbage_price的类型为:[%T],Chinese_cabbage_price的值为:%.2f\n", Chinese_cabbage_price, Chinese_cabbage_price)
	fmt.Printf("weight的类型为:[%T],weight的值为:%d\n", weight, weight)

	/*
		我们知道Chinese_cabbage_price的类型默认是float64类型，weight的类型默认是int类型。

		Go语言和Java不同,Go语言不会进行隐式向上转换类型，无论是向上还是向下转型都需要手动执行。

		综上所述，我们使用int(Chinese_cabbage_price)显式向下转型，即将Chinese_cabbage_price的类型由float64类型转换为int类型，这样可能会导致数据精度丢失,计算结果如下所示。
	*/
	Total_price := int(Chinese_cabbage_price) * weight
	fmt.Printf("大白菜单价是%.2f元/500g,您购买了%d斤,总价为:[%.d]元人名币\n", Chinese_cabbage_price, weight, Total_price)
}
