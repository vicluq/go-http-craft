package router

import (
	mid "github.com/vicluq/golang-api/shared/middleware"
)

func NewRouter(path string) *Router {
	return &Router{
		basePath:    path,
		routes:      make(routeMap),
		middlewares: make([]mid.Middleware, 0),
	}
}
