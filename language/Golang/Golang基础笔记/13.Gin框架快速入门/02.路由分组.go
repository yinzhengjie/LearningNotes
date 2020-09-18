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

	/**
	路由分组:
		在大型项目中，会经常使用到路由分组技术。
		路由分组有点类似于Django创建各种app，其目的就是将项目有组织的划分成多个模块。
	*/
	//定义group1路由组
	group1 := router.Group("group1")
	{
		group1.GET("/login", func(context *gin.Context) {
			context.String(http.StatusOK, "<h1>Login successful</h1>")
		})
	}

	//定义group2路由组
	group2 := router.Group("group2")
	{
		group2.GET("/logout", func(context *gin.Context) {
			context.String(http.StatusOK, "<h3>Logout</h3>")
		})
	}

	//启动路由并指定监听的地址及端口，若不指定默认监听0.0.0.0:8080
	router.Run("127.0.0.1:9000")
}
