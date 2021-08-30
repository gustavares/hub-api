package http

import (
	"github.com/gustavares/hub-api/server/http/router"
	"github.com/valyala/fasthttp"
)

func HealthcheckRoutes() []router.Route {
	return []router.Route{
		{
			Path:    "/healthcheck",
			Method:  fasthttp.MethodGet,
			Handler: HealthcheckHandler(),
		},
		{
			Path:    "/readiness",
			Method:  fasthttp.MethodGet,
			Handler: ReadinessHandler(),
		},
	}
}
