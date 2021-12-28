package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine is the uni handler for all requests
type Engine struct{}

// type HandlerFunc func(ResponseWriter, *Request)
// 参数 ResponseWriter 可以构造针对请求 Request 的响应。
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) ()  {
	// 路由
	switch r.URL.Path {
	case "/":
		_, err := fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
		if err != nil {
			log.Fatal(err.Error())
		}
	case "/hello":
		for k,v := range r.Header {
			_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	default:
		_, err := fmt.Fprintf(w, "404 NOT FOUND: %s \n", r.URL.Path)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func main(){
	// Engine type 实现了接口 Handler，就可以作为实例传入ListenAndServe中了
	engine := new(Engine)
	// ListenAndServe listens on the TCP network address addr and then calls
	// Serve with handler to handle requests on incoming connections.
	// 参数handler是一个接口，需要实现方法 ServeHTTP() ，也就是说，只要传入任何实现了 ServerHTTP 接口的实例，所有的HTTP请求，就都交给该实例处理了。
	// engine实例拦截了所有的HTTP请求，请求有了统一的控制入口，就可以在实例中自定义路由、中间件、日志、异常处理等
	err := http.ListenAndServe(":80", engine)
	if err != nil {
		log.Fatal(err.Error())
	}
}
