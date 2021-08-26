package routes

import (
	"github.com/gustavares/hub-api/server/http"
	"github.com/gustavares/hub-api/server/http/router"
	"github.com/valyala/fasthttp"
)

func HealthcheckRoutes() []router.Route {
	return []router.Route{
		{
			Path:    "/healthcheck",
			Method:  fasthttp.MethodGet,
			Handler: http.HealthcheckHandler(),
		},
		{
			Path:    "/readiness",
			Method:  fasthttp.MethodGet,
			Handler: http.ReadinessHandler(),
		},
	}
}
