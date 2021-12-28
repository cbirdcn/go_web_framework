package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {}
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	// ListenAndServe listens on the TCP network address addr and then calls
	// Serve with handler to handle requests on incoming connections.
	// 参数handler是一个接口，需要实现方法 ServeHTTP() ，也就是说，只要传入任何实现了 ServerHTTP 接口的实例，所有的HTTP请求，就都交给了该实例处理了。
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// type HandlerFunc func(ResponseWriter, *Request)
func indexHandler(w http.ResponseWriter, r *http.Request){
	// Fprintf formats according to a format specifier and writes to w.
	// It returns the number of bytes written and any write error encountered.
	_, err := fmt.Fprintf(w, "URL.Path = %q \n", r.URL.Path)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	for k,v := range r.Header {
		_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
