package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"sierra/common/maingroup"
	"sierra/services/sierra/internal/server/internal/router"
)

func StartBlocking() {
	logrus.Infof("Starting server")
	defer logrus.Infof("Finished gracefully shutting down")

	ctx := context.Background()

	mainGroup := maingroup.New(ctx)

	mainGroup.Go(router.ListenAndServe)

	logrus.Panic(mainGroup.Wait())
}
