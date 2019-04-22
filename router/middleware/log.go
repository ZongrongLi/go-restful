package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/lexkong/log"
	"github.com/tiancai110a/go-rpc/server"
	"github.com/tiancai110a/go-rpc/share"
)

func Logging(rw *http.ResponseWriter, r *http.Request, c *server.Middleware) {

	start := time.Now().UTC()
	path := r.URL.Path

	reg := regexp.MustCompile("(/v1/user|/login)")
	if !reg.MatchString(path) {
		return
	}

	// Skip for the health check requests.
	if path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
		return
	}

	// Read the Body content
	var bodyBytes []byte

	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}

	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// The basic informations.
	method := r.Method
	ip := share.LocalIpV4()

	c.Next(rw, r)

	// Calculates the latency.
	end := time.Now().UTC()
	latency := end.Sub(start)

	code, message := -1, ""

	log.Infof("%-13s | %-12s | %s %s | {code: %d, message: %s}", latency, ip, method, path, code, message)
}
