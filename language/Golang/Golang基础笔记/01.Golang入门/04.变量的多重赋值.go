package main

import (
	"fmt"
)

func main() {

	name := "Jason Yin"
	age := 18
	blog := "https://www.cnblogs.com/yinzhengjie/"

	fmt.Printf("变量name的数据类型是:%T,其对应的值是:%s\n", name, name)
	fmt.Printf("变量age的数据类型是:%T,其对应的值是:%d\n", age, age)
	fmt.Printf("变量blog的数据类型是:%T,其对应的值是:%s\n", blog, blog)

	fmt.Println("**********我是分割线**********")

	/*
		上面的name,age,blog使用三行代码进行自动推导类型赋值，但是写起来相对来说比较复杂。

		我们可以使用多重赋值的方式进行赋值，只需要一行就可以搞定，语法格式如下:
			变量名称1,变量名称2,变量名称3[,...,变量名称n] := 对应变量名称1的值,对应变量名称2的值,对应变量名称3的值[,...,对应变量名称n的值]
	*/
	name2, age2, blog2 := "尹正杰", 27, "https://www.cnblogs.com/yinzhengjie2020"
	fmt.Printf("变量name2的数据类型是:%T,其对应的值是:%s\n", name2, name2)
	fmt.Printf("变量age2的数据类型是:%T,其对应的值是:%d\n", age2, age2)
	fmt.Printf("变量blog2的数据类型是:%T,其对应的值是:%s\n", blog2, blog2)

}
