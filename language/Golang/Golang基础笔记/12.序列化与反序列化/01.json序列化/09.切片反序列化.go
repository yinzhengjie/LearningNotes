package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s1 := `[{"name":"王昭君","role":"中单"},{"name":"李白","role":"打野"}]`

	var slice []map[string]interface{}
	fmt.Println("反序列化之前：slice =", slice)

	/**
	实现思路与前面两种的实现完全一致，这里不再赘述。

	温馨提示:
		反序列化json字符串时，务必确保反序列化传出的数据类型，与之前序列化的数据类型完全一致。
	*/
	err := json.Unmarshal([]byte(s1), &slice)
	if err != nil {
		fmt.Println("反序列化失败,错误原因: ", err)
		return
	}

	fmt.Println("反序列化之后：slice =", slice)
}
