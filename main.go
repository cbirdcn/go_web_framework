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

	// 区分协议类型，第二个参数是HandlerFunc。每个路由pattern可以有自己的handlerFunc
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
		if err != nil {
			log.Fatal(err.Error())
		}
	})

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	})

	// 框架启动http服务
	err := r.Run(":80")
	if err != nil {
		log.Fatal(err.Error())
	}
}
