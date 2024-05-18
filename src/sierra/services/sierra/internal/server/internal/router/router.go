package router

import (
	"context"
	"github.com/luminanceaudio/sierra/src/sierra/common/api"
	"github.com/luminanceaudio/sierra/src/sierra/common/config"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/server/internal/router/routes/audio"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/server/internal/router/routes/samples"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/server/internal/router/routes/shutdown"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/server/internal/router/routes/sources"
	"github.com/sirupsen/logrus"
)

func ListenAndServe(ctx context.Context) error {
	router := api.NewRouter()

	api.RouteGET(router, api.Public, "/api/v1/app", api.SuccessEndpoint)
	api.RoutePOST(router, api.Public, "/api/v1/app/shutdown", shutdown.Shutdown)

	// Sources
	api.RouteGET(router, api.Public, "/api/v1/app/source", sources.GetSources)
	api.RoutePOST(router, api.Public, "/api/v1/app/source", sources.CreateSource)
	api.RoutePOST(router, api.Public, "/api/v1/app/source/delete", sources.DeleteSource)

	// Samples
	api.RoutePOST(router, api.Public, "/api/v1/app/sample/query", samples.QuerySamples)

	api.RoutePOST(router, api.Public, "/api/v1/app/sample/query/count", samples.QuerySamplesCount)
	api.RoutePOST(router, api.Public, "/api/v1/app/sample/query/count/:query", samples.QuerySamplesCount)

	api.RouteGET(router, api.Public, "/api/v1/app/sample/load/*uri", audio.Load)

	logrus.WithField("port", config.InternalIngressPort).Info("Starting HTTP server")
	return router.ListenAndServe(":" + config.InternalIngressPort)
}
