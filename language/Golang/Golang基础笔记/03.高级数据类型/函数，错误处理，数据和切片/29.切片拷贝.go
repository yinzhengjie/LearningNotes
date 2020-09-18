package main

import (
	"fmt"
)

func main() {

	/*
	   Go语言的内置函数copy()可以将一个切片复制到另一个切片。
	   如果加入的两个数组切片不一样，就会按其中较小的那个数组切片的元素个数进行复制。
	*/
	src := []int32{'尹', '正', '杰', '到', '此', '一', '游'}
	fmt.Printf("src的数据为:[%s],长度为:[%d],容量为:[%d],内存地址为:[%p]\n", string(src), len(src), cap(src), src)

	//创建一个长度为5的元素大小的切片
	dest := make([]rune, 5)

	//切片拷贝，会根据较小的切片进行拷贝,由于dest的容量为5，而src的容量为7，将src的元素拷贝到dest时会按照容量较小的(即dest的容量)来进行拷贝
	copy(dest, src)
	fmt.Printf("dest的数据为:[%s],长度为:[%d],容量为:[%d],内存地址为:[%p]\n", string(dest), len(dest), cap(dest), dest)

}
