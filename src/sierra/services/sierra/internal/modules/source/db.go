package source

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"sierra/common/uri"
	"sierra/services/sierra/internal/sierradb"
	source "sierra/services/sierra/internal/sierradb/sierraent/source"
)

func Get(ctx context.Context, uri *uri.URI) (Source, error) {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return nil, err
	}

	src, err := sierraDb.Client.Source.Query().
		Where(source.ID(uri.String())).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	sourceModel, err := New(src.ID)
	if err != nil {
		return nil, err
	}

	return sourceModel, nil
}

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

func Create(ctx context.Context, uri string) error {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return err
	}

	err = sierraDb.Client.Source.Create().
		SetID(uri).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("could not create sample file: %w", err)
	}

	return nil
}
