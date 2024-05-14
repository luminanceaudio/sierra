package source

import (
	"context"
	"fmt"
	"github.com/luminanceaudio/sierra/src/sierra/common/uri"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/client/models"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent"
	source "github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent/source"
	"github.com/sirupsen/logrus"
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

func getAll(ctx context.Context) ([]*sierraent.Source, error) {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return nil, err
	}

	sources, err := sierraDb.Client.Source.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return sources, nil
}

func GetAll(ctx context.Context) ([]Source, error) {
	sources, err := getAll(ctx)
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

func GetAllAsModel(ctx context.Context) ([]*models.Source, error) {
	sources, err := getAll(ctx)
	if err != nil {
		return nil, err
	}

	var sourceList []*models.Source
	for _, src := range sources {
		sourceList = append(sourceList, models.NewSource(src.ID, src.CreateTime))
	}

	return sourceList, nil
}

func Create(ctx context.Context, uri *uri.URI) error {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return err
	}

	err = sierraDb.Client.Source.Create().
		SetID(uri.String()).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("could not create sample file: %w", err)
	}

	return nil
}

func Delete(ctx context.Context, uri *uri.URI) error {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return err
	}

	err = sierraDb.Client.Source.DeleteOneID(uri.String()).Exec(ctx)
	if err != nil {
		return fmt.Errorf("could not create sample file: %w", err)
	}

	return nil
}
