package middleware

import (
	"net/http"
	"time"

	"github.com/tiancai110a/go-rpc/server"
)

func NoCache(rw *http.ResponseWriter, r *http.Request, c *server.Middleware) {

	header := (*rw).Header()
	header.Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	header.Set("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	header.Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next(rw, r)

}

func Options(rw *http.ResponseWriter, r *http.Request, c *server.Middleware) {
	header := (*rw).Header()

	if r.Method != "OPTIONS" {
		c.Next(rw, r)
	} else {
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		header.Set("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		header.Set("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		header.Set("Content-Type", "application/json")
		(*rw).WriteHeader(200)
	}
}

func Secure(rw *http.ResponseWriter, r *http.Request, c *server.Middleware) {
	header := (*rw).Header()

	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("X-Frame-Options", "DENY")
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("X-XSS-Protection", "1; mode=block")
	if r.TLS != nil {
		header.Set("Strict-Transport-Security", "max-age=31536000")
	}
}
