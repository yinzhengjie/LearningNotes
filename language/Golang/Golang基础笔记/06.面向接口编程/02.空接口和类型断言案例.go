package main

import (
	"fmt"
	"reflect"
)

/**
空接口(interface{})不包含任何的方法，正因为如此，所有的类型都实现了空接口，因此空接口可以存储任意类型的数值。
如下所示，我们为空接口起了一个别名.
*/
type MyInterface interface{}

func MyPrint(input MyInterface) {
	/**
	使用断言语法获取传输过来的数据类型，类似于类型强转。
	断言语法格式如下:
		接口类型变量.(断言的类型)
	如果你不确定interface具体是什么类型，那么再断言之前最好先做判断。
	*/
	output, ok := input.(int)
	if ok {
		output = input.(int) + 100 //通过断言语法可以判断数据类型
		fmt.Println(output)
	} else {
		fmt.Println(input)
	}

	inputType := reflect.TypeOf(input) //通过反射也可以判断类型
	fmt.Printf("用户传入的是:[%v],其对应的类型是[%v]\n\n", input, inputType)
}

func main() {
	m1 := true
	MyPrint(m1)

	m2 := "Jason Yin"
	MyPrint(m2)

	m3 := 2020
	MyPrint(m3)

}
