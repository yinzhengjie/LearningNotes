package main

import (
	"fmt"
)

type Animal struct {
	Age int
}

type People struct {
	Animal
	Name   string
	Age    int
	Gender string
}

type IdentityCard struct {
	IdCardNO    int
	Nationality string
	Address     string
	Age         int
}

/*
	此时的Students以及是多重继承
*/
type Students struct {
	IdentityCard
	People //多层继承
	Age    int
	Score  int
}

func main() {
	/**
	如果子类和父类存在同名的属性,那么以就近原则为准
	*/
	s1 := Students{
		Score: 150,
		IdentityCard: IdentityCard{
			IdCardNO:    110105199003072872,
			Nationality: "中华人民共和国",
			Address:     "北京市朝阳区望京SOHO",
			Age:         8,
		},
		People: People{Name: "Jason Yin", Age: 18, Animal: Animal{Age: 20}},
		Age:    27,
	}

	/**
	如果子类和父类存在同名的属性(如果父类还继承了其它类型，我们称之为多层继承)，那么就以就近原则为准;
	但是如果一个子类如果继承自多个父类(我们称之为多重继承),且每个字段中都有相同的字段,此时我们无法直接在子类调用该属性;
	*/
	fmt.Printf("学生的年龄是:[%d]\n", s1.Age)
	s1.Age = 21
	fmt.Printf("学生的年龄是:[%d]\n\n", s1.Age)

	//给People类的Age赋值
	fmt.Printf("People的年龄是:[%d]\n", s1.People.Age)
	s1.People.Age = 5000
	fmt.Printf("People的年龄是:[%d]\n\n", s1.People.Age)

	//给IdentityCard类的Age赋值
	fmt.Printf("IdentityCard的年龄是:[%d]\n", s1.IdentityCard.Age)
	s1.IdentityCard.Age = 80
	fmt.Printf("IdentityCard的年龄是:[%d]\n", s1.IdentityCard.Age)
}
