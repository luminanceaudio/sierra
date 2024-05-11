package router

import (
	"context"
	"sierra/common/api"
	"sierra/common/config"
)

func ListenAndServe(ctx context.Context) error {
	router := api.NewRouter()

	// Register routes
	api.RouteGET(router, api.Public, "/api/v1/app", api.SuccessEndpoint)

	return router.ListenAndServe(":" + config.InternalIngressPort)
}
