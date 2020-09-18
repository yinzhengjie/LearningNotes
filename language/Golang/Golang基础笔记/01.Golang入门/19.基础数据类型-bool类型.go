package main

import (
	"fmt"
)

func main() {

	/*
		Linux内核开源发起人:
			林纳斯·本纳第克特·托瓦兹（Linus Benedict Torvalds, 1969年~ ），著名的电脑程序员。
			Linux内核的发明人及该计划的合作者。托瓦兹利用个人时间及器材创造出了这套当今全球最流行的操作系统（作业系统）内核之一。
			现受聘于开放源代码开发实验室（OSDL：Open Source Development Labs, Inc），全力开发Linux内核。
			该计划开始于1991年.

		Python开源发起人:
			吉多·范罗苏姆（Guido van Rossum，1956年1月31日－） 是一名荷兰计算机程序员。
			他作为 Python 程序设计语言的作者而为人们熟知。
			在 Python 社区，吉多·范罗苏姆被人们认为是“仁慈的独裁者（BDFL）”，意思是他仍然关注 Python 的开发进程，并在必要的时刻做出决定。
			他在 Google 工作，在那里他把一半的时间用来维护 Python 的开发。
			1989年，创立Python语言，1991年初发布第一个公开发行版本.
	*/

	LinusBenedictTorvaldsBirthday := 1969
	GuidoVanRossumBirthday := 1956
	CurrentYear := 2020

	//判断Linux内核开源发起人是否比Python开源发起人年龄大
	res := (CurrentYear - LinusBenedictTorvaldsBirthday) > (CurrentYear - GuidoVanRossumBirthday)

	fmt.Println(res)
	//格式化打印bool值
	fmt.Printf("linux比python年龄大:%t\n", res)
	//格式化打印类型
	fmt.Printf("%T\n", res)
}
