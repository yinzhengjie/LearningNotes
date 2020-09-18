package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	/**
	模型绑定:
		若要将请求主体绑定到结构体中，请使用模型绑定，目前支持JSON、XML、YAML和标准表单值(foo=bar&boo=baz)的绑定。
		需要在绑定的字段上设置tag，比如，绑定格式为json，需要这样设置 json:"fieldname" 。
		你还可以给字段指定特定规则的修饰符，如果一个字段用binding:"required"修饰，并且在绑定时该字段的值为空，那么将返回一个错。
		程序通过tag区分你传递参数的数据格式,从而自动解析相关参数.
	*/
	User   string `form:"user" json:"user" xml:"user"  binding:"required"`
	Passwd string `form:"passwd" json:"passwd" xml:"passwd"  binding:"required"`
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

	router.POST("/login", func(context *gin.Context) {
		//定义接收请求的数据
		var login_user Login

		/**
		Gin还提供了两套绑定方法：
			Must bind:
				Methods - Bind, BindJSON, BindXML, BindQuery, BindYAML Behavior:
					这些方法底层使用MustBindWith，如果存在绑定错误，请求将被以下指令中止 c.AbortWithError(400, err).SetType(ErrorTypeBind)，
					响应状态代码会被设置为400，请求头Content-Type被设置为text/plain; charset=utf-8。
					注意，如果你试图在此之后设置响应代码，将会发出一个警告 [GIN-debug] [WARNING] Headers were already written. Wanted to override status code 400 with 422。
					如果你希望更好地控制行为，请使用ShouldBind相关的方法。

			Should bind
				Methods - ShouldBind, ShouldBindJSON, ShouldBindXML, ShouldBindQuery, ShouldBindYAML Behavior:
					这些方法底层使用 ShouldBindWith，如果存在绑定错误，则返回错误，开发人员可以正确处理请求和错误。
					当我们使用绑定方法时，Gin会根据Content-Type推断出使用哪种绑定器，如果你确定你绑定的是什么，你可以使用MustBindWith或者BindingWith。
		*/
		err := context.ShouldBind(&login_user)
		//如果绑定出错了就将错误信息直接发送给前端页面.
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
		}
		//将结构体绑定后,如果没有报错就可以解析到相应数据，此时我们验证用户名和密码,验证成功返回200状态码，验证失败返回401状态码
		if login_user.User == "yinzhengjie" && login_user.Passwd == "123" {
			context.JSON(http.StatusOK, gin.H{
				"Status": "Login successful\n",
			})
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{
				"Status": "Login failed\n",
			})
		}
	})

	//启动路由并指定监听的地址及端口，若不指定默认监听0.0.0.0:8080
	router.Run("172.30.100.101:9000")

	/**
	使用curl命令进行测试:
		[root@yinzhengjie.com ~]# curl -X POST   http://172.30.100.101:9000/login   -H 'content-type: application/json'   -d '{ "user": "yinzhengjie","passwd":"123"}'

	*/
}
