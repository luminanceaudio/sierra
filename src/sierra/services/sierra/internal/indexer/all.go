package indexer

import (
	"context"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/modules/source"
	"github.com/sirupsen/logrus"
	"time"
)

func IndexAllInterval(ctx context.Context) error {
	err := IndexAll(ctx, true)
	if err != nil {
		logrus.WithError(err).Error("failed to index all during interval (initial run)")
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(30 * time.Second):
			err := IndexAll(ctx, false)
			if err != nil {
				logrus.WithError(err).Error("failed to index all during interval")
			}
		}
	}

}

func IndexAll(ctx context.Context, forceReindex bool) error {
	logrus.Info("Indexing all sources")

	sources, err := source.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, src := range sources {
		err = Singleton().Index(ctx, src, forceReindex, true)
		if err != nil {
			logrus.WithError(err).WithField("sourceUri", src.GetURI().String()).Warn("failed to index source during index all")
			continue
		}
	}

	return nil
}
