package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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

	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 配置8MiB

	router.POST("/upload", func(c *gin.Context) {
		// 多文件
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			//底层采用流拷贝(io.Copy)技术,将上传文件到指定的路径
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	//启动路由并指定监听的地址及端口，若不指定默认监听0.0.0.0:8080
	router.Run("172.30.100.101:9000")

	/**
	使用curl命令测试:
		[root@yinzhengjie.com ~]# curl -X POST http://172.30.100.101:9000/upload   -F "upload[]=@/etc/issue"   -F "upload[]=@/etc/passwd"   -H "Content-Type: multipart/form-data"
	*/
}
