package frame

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// 将http请求的method和pattern(比如/)拼接后， 作为router map 的唯一key
func (r *router) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	engine.router[method+"-"+pattern] = handlerFunc
}