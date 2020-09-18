package main

import "fmt"

func main() {

	/*
		字符(byte)类型是uint8的别名（type为已存在的数据类型起一个别名），在builtin.go文件中可以看到如下的代码:
			type byte = uint8

				ASCII编码表称为美国信息交换码，其结构大致如下:
			0-31:
				表示控制字符，在键盘上找不到对应的按键，但是可以使用转义字符表示
			32-126:
				键盘上可以找到的所有字符
			127:
				删除(del)键

		常见的转义字符如下:
			'\n':
				换行符,对应ASCII编码表的十进制是10
			'\t':
				水平制表符，对应ASCII编码表的十进制是9
			'\v':
				垂直制表符,对应ASCII编码表的十进制是11
			'\r':
				回车,对应ASCII编码表的十进制是13
			'0':
				数字字符0对应ASCII编码表的十进制为48
			'A':
				大写字母字符A对应ASCII编码表的十进制是65
			'a':
				小写字母字符a对应ASCII编码表的十进制是97
			'del':
				删除键对应ASCII编码表的十进制是127
	*/
	var a byte = 'a'

	/*
		采用uint8数据格式打印，最终输出的是字符'a'在ASCII所对应的十进制数字编码值为97
	*/
	fmt.Println(a)

	/*
		%c是一个占位符，表示输出一个字符类型数据。
	*/
	fmt.Printf("a的类型为:[%T],a的值为:[%c]\n", a, a)

	var b = '0'
	fmt.Println(b)
}
