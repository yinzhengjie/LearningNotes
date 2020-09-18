package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	/**
	所有的接口都要由路由来进行管理。
		Gin的路由支持GET,POST,PUT,DELETE,PATCH,HEAD,OPTIONS等请求
		同时还有一个Any函数，可以同时支持以上的所有请求。

	创建路由(router)并引入默认中间件
		在源码中,首先是New一个engine,紧接着通过Use方法传入了Logger()和Recovery()这两个中间件。
		其中 Logger 是对日志进行记录，而 Recovery 是对有 painc时, 进行500的错误处理。

	创建路由(router)无中间件
		router := gin.New()
	*/
	router := gin.Default()

	router.POST("/blog", func(context *gin.Context) {
		//从POST方法获取参数
		user := context.PostForm("user")
		//获取POST方法带默认值参数,如果没有则返回默认值"yinzhengjie"
		passwd := context.DefaultPostForm("passwd", "yinzhengjie")
		//将获取到的数据返回给客户端
		context.JSON(http.StatusOK, gin.H{
			"status": "POST",
			"USER":   user,
			"PASSWD": passwd,
		})
	})

	//启动路由并指定监听的地址及端口，若不指定默认监听0.0.0.0:8080
	router.Run("172.30.100.101:9000")
}
