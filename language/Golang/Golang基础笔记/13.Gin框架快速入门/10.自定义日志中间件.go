package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
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
	router := gin.New()

	//创建一个日志文件
	f, _ := os.Create("gin.log")

	//默认数据写入到终端控制台(os.Stdout),我们需要将日志写到咱们刚刚创建的日志文件中
	gin.DefaultWriter = io.MultiWriter(f)

	//自定义你的日志格式
	logger := func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			//客户端IP
			params.ClientIP,
			//请求时间
			params.TimeStamp.Format(time.RFC1123),
			//请求方法
			params.Method,
			//请求路径
			params.Path,
			//请求协议
			params.Request.Proto,
			//请求的状态码
			params.StatusCode,
			//请求延迟(耗时)
			params.Latency,
			//请求的客户端类型
			params.Request.UserAgent(),
			//请求的错误信息
			params.ErrorMessage,
		)
	}

	//LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
	router.Use(gin.LoggerWithFormatter(logger))

	router.GET("/log", func(context *gin.Context) {
		context.String(http.StatusOK, "自定义日志中间件\n")
	})

	//启动路由并指定监听的地址及端口，若不指定默认监听0.0.0.0:8080
	router.Run("172.30.100.101:9000")

}
