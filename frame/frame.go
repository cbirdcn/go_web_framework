package frame

import (
	"fmt"
	"log"
	"net/http"
)

// 自定义的"请求-返回"处理函数类型
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 的router用map保存从路由到 HandlerFunc 的映射
// engine Engine作为ListenAndServe()的第二个参数， 需要实现Handler接口的ServeHttp()方法才能生成实例engine并调用方法
type Engine struct {
	router map[string]HandlerFunc
}

// New is the constructor of frame.Engine
// 注意：只能实例化1次
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 将http请求的method和pattern(比如/)拼接后， 作为router map 的唯一key
func (engine *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	engine.router[method+"-"+pattern] = handlerFunc
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	// 只有engine 实现了ServeHTTP时，才可以作为第二个参数
	return http.ListenAndServe(addr, engine)
}

// 实现接口 Handler 的方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		// 设置返回状态码
		w.WriteHeader(http.StatusNotFound)
		_, err := fmt.Fprintf(w, "404 NOT FOUND: %s \n", r.URL.Path)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

