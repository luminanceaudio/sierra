package router

import (
	"context"
	"github.com/sirupsen/logrus"
	"sierra/common/api"
	"sierra/common/config"
	"sierra/services/sierra/internal/server/internal/router/routes/audio"
	"sierra/services/sierra/internal/server/internal/router/routes/samples"
	"sierra/services/sierra/internal/server/internal/router/routes/shutdown"
	"sierra/services/sierra/internal/server/internal/router/routes/sources"
)

func ListenAndServe(ctx context.Context) error {
	router := api.NewRouter()

	api.RouteGET(router, api.Public, "/api/v1/app", api.SuccessEndpoint)
	api.RoutePOST(router, api.Public, "/api/v1/app/shutdown", shutdown.Shutdown)

	// Sources
	api.RouteGET(router, api.Public, "/api/v1/app/source", sources.GetSources)
	api.RoutePOST(router, api.Public, "/api/v1/app/source", sources.CreateSource)
	api.RoutePOST(router, api.Public, "/api/v1/app/source/delete", sources.DeleteSource)

	api.RouteGET(router, api.Public, "/api/v1/app/sample", samples.GetSamples)
	api.RouteGET(router, api.Public, "/api/v1/app/sample/load/*uri", audio.Load)
	api.RouteGET(router, api.Public, "/api/v1/app/sample/waveform/*uri", audio.GetWaveform)

	logrus.WithField("port", config.InternalIngressPort).Info("Starting HTTP server")
	return router.ListenAndServe(":" + config.InternalIngressPort)
}
