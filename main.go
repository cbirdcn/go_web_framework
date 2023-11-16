package main

import (
	"go_web_framework/frame" // go mod init go_web_framework后，从module/frame导入包
	"fmt"
	"log"
	"net/http"
)

func main(){
	// 调用框架实例化
	r := frame.New()

	// 要构造一个完整的响应，需要考虑消息头(Header)和消息体(Body)，而 Header 包含了状态码(StatusCode)，消息类型(ContentType)等几乎每次请求都需要设置的信息。
	// 因此，如果不进行有效的封装，那么框架的用户将需要写大量重复，繁杂的代码，而且容易出错。
	// Handler的参数从HandlerFunc变成了frame.Context，提供了Query/PostForm参数等功能。
	// frame.Context封装了HTML/String/JSON等函数，能够快速构造HTTP响应。
	// 封装*http.Request和http.ResponseWriter的方法，简化相关接口的调用，只是设计 Context 的原因之一。对于框架来说，还需要支撑额外的功能。
	// 例如，将来解析动态路由/hello/:name，参数:name的值放在哪呢？
	// 再比如，框架需要支持中间件，那中间件产生的信息放在哪呢？Context 随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由 Context 承载。
	// 因此，设计 Context 结构，扩展性和复杂性留在了内部，而对外简化了接口。路由的处理函数，以及将要实现的中间件，参数都统一使用 Context 实例， Context 就像一次会话的百宝箱，可以找到任何东西。
	r.GET("/", func(c *frame.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *frame.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		for k, v := range r.Header {
			_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	})

	r.POST("/login", func(c *frame.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	// 框架启动http服务
	err := r.Run(":80")
	if err != nil {
		log.Fatal(err.Error())
	}
}
