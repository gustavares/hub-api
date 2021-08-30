package router

import (
	httprouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
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

func New(routes *[]Route) Router {
	r := &Router{
		Router: httprouter.New(),
	}
	r.addRoutes(*routes...)

	return *r
}

func (r Router) addRoutes(routes ...Route) {
	for _, route := range routes {
		r.Router.Handle(route.Method, route.Path, route.Handler)
	}
}
