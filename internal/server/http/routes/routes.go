package routes

import (
	"github.com/gustavares/hub-api/internal/server/http/router"
)

func BuildRoutes() *[]router.Route {
	allRoutes := [][]router.Route{
		HealthcheckRoutes(),
		UserRoutes(),
	}

	appendRoutes := []router.Route{}

	for _, routeGroup := range allRoutes {
		appendRoutes = append(appendRoutes, routeGroup...)
	}

	return &appendRoutes
}
