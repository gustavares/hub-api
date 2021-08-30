package router

import (
	httprouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var (
	WithRoutes = func(routes ...Route) ConfigRouter {
		return func(router *Router) {
			router.AddRoutes(routes...)
		}
	}
)

type Route struct {
	Path    string
	Method  string
	Handler fasthttp.RequestHandler
}
type Router struct {
	Router *httprouter.Router
}

type ConfigRouter func(router *Router)

func New(configs ...ConfigRouter) Router {
	r := &Router{
		Router: httprouter.New(),
	}

	for _, config := range configs {
		config(r)
	}

	return *r
}

func (r Router) AddRoutes(routes ...Route) {
	for _, route := range routes {
		r.Router.Handle(route.Method, route.Path, route.Handler)
	}
}
