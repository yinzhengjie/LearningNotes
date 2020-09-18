package main

import (
	"fmt"
	"net/http"
)

func main() {

	/**
	该URL是咱们自己刚刚编写简单的Web服务器对应的资源。
	*/
	url := "http://127.0.0.1:8080/user"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取数据失败,错误原因: ", err)
		return
	}
	defer resp.Body.Close()

	/**
	获取从服务器端读到数据
	*/
	fmt.Printf("状态: %s\n", resp.Status)
	fmt.Printf("状态码: %v\n", resp.StatusCode)
	fmt.Printf("响应头部: %s\n", resp.Header)
	fmt.Println("响应包体: ", resp.Body)

	/**
	定义切片缓冲区，临时存储读到的数据,并将每次读到的结果拼接到data中
	*/
	buf := make([]byte, 4096)
	var data string

	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("数据读取完毕....")
			break
		}
		if err != nil {
			fmt.Println("数据读取失败,错误原因: ", err)
			return
		}
		data += string(buf[:n])
	}

	fmt.Printf("从服务端获取到的内容是: [%s]\n", data)
}
