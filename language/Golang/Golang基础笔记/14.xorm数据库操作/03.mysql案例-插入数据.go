package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

//结构体字段对应数据库中的表字段
type User struct {
	Id      int64
	Name    string `xorm:"name"`
	Age     int    `xorm:"age"`
	Phone   string `xorm:"phone"`
	Address string `xorm:"address"`
}

func main() {
	/**
	配置连接数据库信息格式如下所示:
		用户名:密码@tcp(数据库服务器地址:端口)/数据库名称?charset=字符集
	*/
	cmd := fmt.Sprintf("jason:yinzhengjie@tcp(172.200.1.254:3306)/golang?charset=utf8mb4")

	//使用上面的配置信息连接数据库,但要指定数据库的类型(我这里写MySQL)
	db_conn, err := xorm.NewEngine("mysql", cmd)
	if err != nil {
		log.Fatal(err)
	}

	//释放资源
	defer db_conn.Close()

	//定义待插入用户的数据
	new_user := User{
		Id:      3,
		Name:    "诡术妖姬",
		Age:     25,
		Phone:   "1000001",
		Address: "艾欧尼亚",
	}

	n, err := db_conn.Insert(new_user)
	fmt.Printf("成功插入了[%d]条数据!\n", n)
}
