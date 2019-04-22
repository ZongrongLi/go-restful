package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/viper"
	"github.com/tiancai110a/go-restful/pkg/errno"
	"github.com/tiancai110a/go-restful/pkg/token"
	"github.com/tiancai110a/go-rpc/server"
)

func AuthMiddleware(rw *http.ResponseWriter, r *http.Request, c *server.Middleware) {

	secret := viper.GetString("jwt_secret")
	auth := r.Header.Get("Authorization")

	if len(auth) == 0 {
		header := (*rw).Header()
		header.Set("Statuscode", strconv.FormatInt(int64(errno.ErrTokenInvalid.Code), 10))
		header.Set("message", errno.ErrTokenInvalid.Message)
		return
	}
	var t string
	fmt.Sscanf(auth, "Bearer %s", &t)

	if _, err := token.Parse(t, secret); err != nil {
		header := (*rw).Header()
		header.Set("Statuscode", strconv.FormatInt(int64(errno.ErrTokenInvalid.Code), 10))
		header.Set("message", errno.ErrTokenInvalid.Message)
		return
	}

	c.Next(rw, r)

}
