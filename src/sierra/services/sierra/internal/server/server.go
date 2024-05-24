package server

import (
	"context"
	"github.com/google/uuid"
	"github.com/luminanceaudio/sierra/src/sierra/common/maingroup"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/indexer"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/modules/collection"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/server/internal/router"
)

func StartBlocking() error {
	ctx := context.Background()

	err := collection.CreateIfNotExist(ctx, uuid.MustParse("fd832e80-5e17-4966-ab9b-7616fafaa63b"), "My samples")
	if err != nil {
		return err
	}

	mainGroup := maingroup.New(ctx)

	mainGroup.Go(router.ListenAndServe)
	mainGroup.Go(indexer.IndexAllInterval)

	return mainGroup.Wait()
}
