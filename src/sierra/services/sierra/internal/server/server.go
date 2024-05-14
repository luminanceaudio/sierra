package server

import (
	"context"
	"github.com/luminanceaudio/sierra/src/sierra/common/maingroup"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/indexer"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/server/internal/router"
)

func StartBlocking() error {
	ctx := context.Background()

	mainGroup := maingroup.New(ctx)

	mainGroup.Go(router.ListenAndServe)
	mainGroup.Go(indexer.IndexAllInterval)

	return mainGroup.Wait()
}
