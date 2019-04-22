package middleware

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	"github.com/tiancai110a/go-rpc/server"
)

func RequestId(rw *http.ResponseWriter, r *http.Request, c *server.Middleware) {

	// Check for incoming header, use it if exists
	requestId := r.Header.Get("X-Request-Id")

	// Create request id with UUID4
	if requestId == "" {
		u4, _ := uuid.NewV4()
		requestId = u4.String()
	}

	// Expose it for use in the application
	r.Header.Set("X-Request-Id", requestId)

	// Set X-Request-Id header
	header := (*rw).Header()
	header.Set("X-Request-Id", requestId)
	c.Next(rw, r)
}
