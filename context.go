package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 给map[string]interface{}起了一个别名gee.H，构建JSON数据时，显得更简洁。
type H map[string]interface{}

// Context目前只包含了 http.ResponseWriter 和 *http.Request，另外提供了对 Method 和 Path 这两个常用属性的直接访问。
type Context struct {
	Writer http.ResponseWriter
	Req *http.Request

	Path string
	Method string

	StatusCode int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context{
	return &Context{
		Writer:     w,
		Req:        r,
		Path:       r.URL.Path,
		Method:     r.Method,
	}
}

// 提供了访问Query和PostForm参数的方法。
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 提供了快速构造String/Data/JSON/HTML响应的方法。
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}