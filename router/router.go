package router

import (
	"github.com/tiancai110a/go-restful/router/middleware"
	"github.com/tiancai110a/go-restful/view"
	"github.com/tiancai110a/go-rpc/server"
	"github.com/tiancai110a/go-rpc/service"
)

// Load loads the middlewares, routes, handlers.
func Load(s server.RPCServer) {
	// Middlewares.

	s.Use(middleware.TestMiddleware1)
	s.Use(middleware.TestMiddleware2)
	s.Use(middleware.TestMiddleware3)

	// The health check handlers
	svcd := s.Group(service.GET, "/view")
	{
		svcd.Route("/health", view.HealthCheck)
		svcd.Route("/disk", view.DiskCheck)
		svcd.Route("/cpu", view.CPUCheck)
		svcd.Route("/ram", view.RAMCheck)
	}

	return
}
