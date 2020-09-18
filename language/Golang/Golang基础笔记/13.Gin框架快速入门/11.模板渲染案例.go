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
		router := gin.Default()
		在源码中,首先是New一个engine,紧接着通过Use方法传入了Logger()和Recovery()这两个中间件。
		其中 Logger 是对日志进行记录，而 Recovery 是对有 painc时, 进行500的错误处理。

	创建路由(router)无中间件
		router := gin.New()
	*/
	router := gin.Default()

	//加载模版文件
	router.LoadHTMLGlob("./templates/*")

	//加载静态资源
	router.Static("/statics", "./statics")

	//定义GET方法
	router.GET("/index", func(context *gin.Context) {
		/**
		http.StatusOK:
			响应的状态码,对应常量200.
		"index.tmpl":
			前端模板的名字。

		温馨提示:
			以"*.tmpl"结尾的文件是Golang针对前端页面进行定义的格式,事实上它也是一种模板引擎(比如逻辑运算,判断,循环，变量,过滤等功能)。
		*/
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			//改变量在"index.tmpl"中是有定义的，可以通过相同变量名拿到对应的数据哟~
			"title": "Hello World !",
		})
	})

	//启动路由并指定监听的地址及端口，若不指定默认监听0.0.0.0:8080
	router.Run("172.30.100.101:9000")
}
