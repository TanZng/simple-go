package routes

import (
	"simple-go/cmd/server/api"
)

type routeAbstract interface {
	RegisterRoutes(*api.ApiService)
}

type routeService struct {
	api   *api.ApiService
	route routeAbstract
}

func New(api *api.ApiService, routes routeAbstract) routeService {
	return routeService{
		api:   api,
		route: routes,
	}
}

func (r routeService) Register() {
	r.route.RegisterRoutes(r.api)
}
