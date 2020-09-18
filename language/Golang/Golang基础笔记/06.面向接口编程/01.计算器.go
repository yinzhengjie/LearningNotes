package main

import (
	"fmt"
)

//实现面向对象版本包含加减法的计算器
type Parents struct {
	x int
	y int
}

//实现加法类
type Addition struct {
	Parents
}

//实现减法类
type Subtraction struct {
	Parents
}

//实现乘法类
type multiplication struct {
	Parents
}

//实现除法类
type Division struct {
	Parents
}

func (this *Addition) Operation() int {
	return this.x + this.y
}

func (this *Subtraction) Operation() int {
	return this.x - this.y
}

func (this *multiplication) Operation() int {
	return this.x * this.y
}

func (this *Division) Operation() int {
	return this.x / this.y
}

/**
实现接口版本包含加减法的计算器

接口就是一种规范标准,接口中不做函数实现，只做定义函数格式

面向接口编程(也称为面向协议编程)降低了代码的耦合度,方便后期代码的维护和扩充,这种实现方法我们称之为多态。

多态三要素:
	1>.父类是接口;
	2>.子类实现所有的接口中定义的函数;
	3>.有一个父类接口对应子类对象指针;
*/
type MyCalculator interface {
	Operation() int //实现接口的结构体中必须包含Operation函数名且返回值为int类型
}

func Calculation(c MyCalculator) int {
	return c.Operation()
}

func main() {
	//调用加法
	a := Addition{Parents{100, 20}}
	sum := a.Operation()
	fmt.Println(sum)

	//调用减法
	b := Subtraction{Parents{100, 20}}
	sub := b.Operation()
	fmt.Println(sub)

	//调用乘法
	c := multiplication{Parents{100, 20}}
	mul := c.Operation()
	fmt.Println(mul)

	//调用除法
	d := Division{Parents{100, 20}}
	div := d.Operation()
	fmt.Println(div)

	fmt.Println("===== 我是分割线 =====")

	//调用接口,需要传入对象指针，相比上面面向对象的方法来说,接口表现了面向接口三要素中的多态特征。
	fmt.Println(Calculation(&a))
	fmt.Println(Calculation(&b))
	fmt.Println(Calculation(&c))
	fmt.Println(Calculation(&d))
}
