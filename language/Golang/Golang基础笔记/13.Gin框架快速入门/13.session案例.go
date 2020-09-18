package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	//定义加密
	store := cookie.NewStore([]byte("secret"))

	//绑定session中间件
	router.Use(sessions.Sessions("mysession", store))

	//定义GET方法
	router.GET("/session", func(context *gin.Context) {
		//初始化session对象
		session := sessions.Default(context)

		//如果浏览器第一次访问返回状态码401，第二次访问则返回状态码200
		if session.Get("user") != "yinzhengjie" {
			session.Set("user", "yinzhengjie")
			session.Save()
			context.JSON(http.StatusUnauthorized, gin.H{"user": session.Get("user")})
		} else {
			context.String(http.StatusOK, "Successful second visit")
		}

	})

	//启动路由并指定监听的地址及端口，若不指定默认监听0.0.0.0:8080
	router.Run("172.30.100.101:9000")

	/**
	  	测试工具建议使用浏览器访问"http://172.30.100.101:9000/session“,不推荐使用curl命令。
		因为curl工具无法缓存,浏览器是由缓存的可以很明显看到测试效果。
	*/
}
