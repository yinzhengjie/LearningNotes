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

	//定义待更新的字段内容
	user := User{Age: 27}

	//定义删除name字段为"诡术妖姬"的数据
	n, err := db_conn.Where("id = ?", 1).Update(user)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("成功更新了[%d]条数据!\n", n)
}
