package server

import (
	"context"
	"sierra/common/maingroup"
	"sierra/services/sierra/internal/server/internal/router"
)

func StartBlocking() error {
	ctx := context.Background()

	mainGroup := maingroup.New(ctx)

	mainGroup.Go(router.ListenAndServe)

	return mainGroup.Wait()
}
