package source

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"sierra/services/sierra/internal/sierradb"
)

func GetAll(ctx context.Context) ([]Source, error) {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return nil, err
	}

	sources, err := sierraDb.Client.Source.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var sourceList []Source
	for _, src := range sources {
		sourceModel, err := New(src.ID)
		if err != nil {
			logrus.WithError(err).Error("failed to get source model")
			continue
		}

		sourceList = append(sourceList, sourceModel)
	}

	return sourceList, nil
}

func CreateLocal(ctx context.Context, path string) error {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("file://%s", path)

	err = sierraDb.Client.Source.Create().
		SetID(uri).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("could not create sample file: %w", err)
	}

	return nil
}
