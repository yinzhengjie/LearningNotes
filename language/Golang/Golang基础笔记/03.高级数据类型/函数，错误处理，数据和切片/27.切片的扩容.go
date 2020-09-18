package main

import (
	"fmt"
)

func main() {

	/*
	   切片的扩容:
	       切片的长度时不固定的，可以向已经定义的切片中追加数据。
	       通过append函数可以在原切片的尾部添加元素

	   切片的结构如下所示，其中unsafe.Pointer对应的类型为:
	       type slice struct {
	           array unsafe.Pointer
	           len   int
	           cap   int
	       }
	       接下来我们来刨析一下slice结构体:
	           array:
	               存储的是数组指针,它指向了数组的内存地址。
	               其类型"unsafe.Pointer"对应“type Pointer *ArbitraryType”,而"ArbitraryType"对应的是"int"，而int在64为操作系统上来讲就是int64，对应8字节。
	           len:
	               int类型对应的是8字节
	           cap:
	               int类型对应的是8字节
	*/
	s1 := []int{100, 200, 300} //使用自动推导类型定义切片
	fmt.Printf("s1的长度为:%d,容量为:%d,内存地址为:%p\n", len(s1), cap(s1), s1)

	/*
	   向slice1切片中添加(扩容)了5个元素，尽管slice1之前的容量是3，但依旧是可以往里面添加数据的哟~而且并不会改变slice1的内存地址。
	*/
	s1 = append(s1, 100)
	fmt.Printf("s1的长度为:%d,容量为:%d,内存地址为:%p\n", len(s1), cap(s1), s1)

	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Printf("s1的长度为:%d,容量为:%d,内存地址为:%p\n", len(s1), cap(s1), s1)

	/*
	   如果切片扩容超过切片的容量，那么Go语言会该切片重新申请一个内存空间，把原来切片中的数据拷贝到新的地址空间，因此在生产环境中建议大家不要将切片的容量设置过小。

	   切片的扩容规则(参考源码文件"runtime/slice.go"的"func growslice(et *_type, old slice, cap int) slice"函数):
	   　　　　如果长度和容量相等分为下面三种情况:
	   　　　　　　1>.长度和容量小于等于896时:
	   　　　　　　　　再次为切片添加数据时，切片会自动扩容，扩容大小为上一次的2倍。
	   　　　　　　2>.长度和容量大于896小于1024时:
	   　　　　　　　　再次为切片添加数据时，切片会自动扩容，扩容大小固定为2048。
	   　　　　　　3>.长度和容量大于等于1024时:
	   　　　　　　　　再次为切片添加数据时，切片会自动扩容扩容为上一次的四分之一(1/4)。

	*/
	s2 := make([]int, 896, 896)
	fmt.Printf("s2的长度为:%d,容量为:%d,内存地址为:%p\n", len(s2), cap(s2), s2)
	s2 = append(s2, 1) //
	fmt.Printf("s2的长度为:%d,容量为:%d,内存地址为:%p\n", len(s2), cap(s2), s2)

	s3 := make([]int, 897, 897)
	fmt.Printf("s3的长度为:%d,容量为:%d,内存地址为:%p\n", len(s3), cap(s3), s3)
	s3 = append(s3, 2) //
	fmt.Printf("s3的长度为:%d,容量为:%d,内存地址为:%p\n", len(s3), cap(s3), s3)

	s4 := make([]int, 1023, 1023)
	fmt.Printf("s4的长度为:%d,容量为:%d,内存地址为:%p\n", len(s4), cap(s4), s4)
	s4 = append(s4, 2) //
	fmt.Printf("s4的长度为:%d,容量为:%d,内存地址为:%p\n", len(s4), cap(s4), s4)

	s5 := make([]int, 1024, 1024)
	fmt.Printf("s5的长度为:%d,容量为:%d,内存地址为:%p\n", len(s5), cap(s5), s5)
	s5 = append(s5, 2) //
	fmt.Printf("s5的长度为:%d,容量为:%d,内存地址为:%p\n", len(s5), cap(s5), s5)

	s6 := make([]int, 1025, 1025)
	fmt.Printf("s6的长度为:%d,容量为:%d,内存地址为:%p\n", len(s6), cap(s6), s6)
	s6 = append(s6, 2)
	fmt.Printf("s6的长度为:%d,容量为:%d,内存地址为:%p\n", len(s6), cap(s6), s6)
}
