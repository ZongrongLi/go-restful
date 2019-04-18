package middleware

import (
	"fmt"
	"net/http"

	"github.com/tiancai110a/go-rpc/server"
)

func TestMiddleware1(rw *http.ResponseWriter, r *http.Request, c *server.Middleware) {

	fmt.Println("before===testMiddlewarec1")
	c.Next(nil, nil)

	fmt.Println("after===testMiddlewarec1")
}

func TestMiddleware2(rw *http.ResponseWriter, r *http.Request, c *server.Middleware) {
	fmt.Println("before===testMiddlewarec2")
	c.Next(nil, nil)

	fmt.Println("after===testMiddlewarec2")
}

func TestMiddleware3(rw *http.ResponseWriter, r *http.Request, c *server.Middleware) {
	fmt.Println("before===testMiddlewarec3")
	c.Next(nil, nil)
	fmt.Println("after===testMiddlewarec3")
}
