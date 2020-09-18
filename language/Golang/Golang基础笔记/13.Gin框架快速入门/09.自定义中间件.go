package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
自定义一个中间件功能:
	返回的包头(header)信息有咱们自定义的包头信息。
*/
func ResponseHeaders() gin.HandlerFunc {
	return func(context *gin.Context) {
		//自定义包头信息
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CRSF-Token,Authorization,Token")
		context.Header("Access-Control-Allow-Methods", "POST,GET,DELETE,OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		//使用"context.Next()"表示继续调用其它的内置中间件，也可以立即终止调用其它的中间件使用"context.Abort()"
		context.Next()

	}
}

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

	//绑定咱们自己定义的中间件
	router.Use(ResponseHeaders())

	router.GET("/middle", func(context *gin.Context) {
		context.String(http.StatusOK, "Response OK\n")
	})

	//启动路由并指定监听的地址及端口，若不指定默认监听0.0.0.0:8080
	router.Run("172.30.100.101:9000")

	/**
	使用curl命令测试:
		[root@yinzhengjie.com ~]# curl -v   http://172.30.100.101:9000/middle
	*/
}
