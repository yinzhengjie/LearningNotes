package main

import (
	"fmt"
	"time"
)

type Vehicle struct {
	Brand string
	Wheel byte
}

type Car struct {
	Vehicle
	Colour string
}

type Driver struct {
	*Car
	DrivingTime time.Time
}

func main() {
	/**
	对象指针匿名组合的第一种初始化方式:
		定义时直接初始化赋值。
	*/
	d1 := Driver{&Car{
		Vehicle: Vehicle{
			Brand: "丰田",
			Wheel: 4,
		},
		Colour: "红色",
	}, time.Now(),
	}
	//打印结构体的详细信息,注意观察指针对象
	fmt.Printf("%+v\n", d1)
	//我们可以直接调用对象的属性
	fmt.Printf("品牌:%s,颜色:%s\n", d1.Brand, d1.Colour)
	fmt.Printf("驾驶时间:%+v\n\n", d1.DrivingTime)
	time.Sleep(1000000000 * 3)

	/**
	对象指针匿名组合的第二种初始化方式:
		先声明，再赋值
	温馨提示:
		遇到指针的情况一定要避免空(nil)指针,未初始化的指针默认值是nil,可以考虑使用new函数解决。
	*/
	var d2 Driver
	/**
	由于Driver结构体中有一个对象指针匿名组合Car,因此我们需要使用new函数申请空间。
	*/
	d2.Car = new(Car)
	d2.Brand = "奔驰"
	d2.Colour = "黄色"
	d2.DrivingTime = time.Now()
	fmt.Printf("%+v\n", d2)
	fmt.Printf("品牌:%s,颜色:%s\n", d2.Brand, d2.Colour)
	fmt.Printf("驾驶时间:%+v\n", d1.DrivingTime)
}
