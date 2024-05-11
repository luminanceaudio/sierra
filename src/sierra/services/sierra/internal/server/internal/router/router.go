package router

import (
	"context"
	"github.com/sirupsen/logrus"
	"sierra/common/api"
	"sierra/common/config"
	"sierra/services/sierra/internal/server/internal/router/routes/samples"
	"sierra/services/sierra/internal/server/internal/router/routes/shutdown"
)

func ListenAndServe(ctx context.Context) error {
	router := api.NewRouter()

	// Register routes
	api.RouteGET(router, api.Public, "/api/v1/app", api.SuccessEndpoint)
	api.RoutePOST(router, api.Public, "/api/v1/app/shutdown", shutdown.Shutdown)
	api.RoutePOST(router, api.Public, "/api/v1/app/samples", samples.GetSamples)

	logrus.WithField("port", config.InternalIngressPort).Info("Starting HTTP server")
	return router.ListenAndServe(":" + config.InternalIngressPort)
}
