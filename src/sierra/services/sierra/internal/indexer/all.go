package indexer

import (
	"context"
	"github.com/sirupsen/logrus"
	"sierra/services/sierra/internal/modules/source"
	"time"
)

func IndexAllInterval(ctx context.Context) error {
	err := IndexAll(ctx)
	if err != nil {
		logrus.WithError(err).Error("failed to index all during interval (initial run)")
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(30 * time.Second):
			err := IndexAll(ctx)
			if err != nil {
				logrus.WithError(err).Error("failed to index all during interval")
			}
		}
	}

}

func IndexAll(ctx context.Context) error {
	logrus.Info("Indexing all sources")

	sources, err := source.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, src := range sources {
		err = Singleton().Index(ctx, src, false, true)
		if err != nil {
			logrus.WithError(err).WithField("sourceUri", src.GetURI().String()).Warn("failed to index source during index all")
			continue
		}
	}

	return nil
}
