package sourcesample

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"sierra/common/sha256"
	"sierra/common/uri"
	"sierra/services/sierra/client/models"
	"sierra/services/sierra/internal/sierradb"
	sample2 "sierra/services/sierra/internal/sierradb/sierraent/sample"
	"sierra/services/sierra/internal/sierradb/sierraent/source"
	"sierra/services/sierra/internal/sierradb/sierraent/sourcesample"
)

func Get(ctx context.Context, uri *uri.URI) (*models.Sample, error) {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return nil, err
	}

	sourceSample, err := sierraDb.Client.SourceSample.Query().
		WithSample().
		WithSource().
		Where(sourcesample.ID(uri.String())).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	sample := models.NewSample(
		sourceSample.ID,
		sourceSample.Edges.Sample.ID,
		sourceSample.Edges.Sample.Format,
		sourceSample.Edges.Source.ID,
	)

	return sample, nil
}

func GetBySampleSha256(ctx context.Context, sha256 sha256.Sha256) (*models.Sample, error) {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return nil, err
	}

	sourceSample, err := sierraDb.Client.SourceSample.Query().
		WithSample().
		WithSource().
		Where(sourcesample.HasSampleWith(sample2.ID(sha256))).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	sample := models.NewSample(
		sourceSample.ID,
		sourceSample.Edges.Sample.ID,
		sourceSample.Edges.Sample.Format,
		sourceSample.Edges.Source.ID,
	)

	return sample, nil
}

func GetAll(ctx context.Context) ([]*models.Sample, error) {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return nil, err
	}

	sourceSamples, err := sierraDb.Client.SourceSample.Query().WithSource().WithSample().All(ctx)
	if err != nil {
		return nil, err
	}

	var results []*models.Sample
	for _, sample := range sourceSamples {
		if sample.Edges.Source == nil {
			logrus.Error("source is nil")
			continue
		}

		if sample.Edges.Sample == nil {
			logrus.Error("sample is nil")
			continue
		}

		results = append(results, models.NewSample(
			sample.ID,
			sample.Edges.Sample.ID,
			sample.Edges.Sample.Format,
			sample.Edges.Source.ID,
		))
	}
	return results, nil
}

func Upsert(ctx context.Context, sha256Str string, sourceUri, fileUri *uri.URI) error {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return err
	}

	err = sierraDb.Client.SourceSample.Create().
		SetID(fileUri.String()).
		SetSampleID(sha256Str).
		SetSourceID(sourceUri.String()).
		OnConflictColumns(sourcesample.FieldID).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("could not create source sample: %w", err)
	}

	return nil
}

func DeleteBySource(ctx context.Context, sourceUri *uri.URI) error {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return err
	}

	deleted, err := sierraDb.Client.SourceSample.Delete().
		Where(sourcesample.HasSourceWith(source.ID(sourceUri.String()))).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("could not delete sourcesamples: %w", err)
	}

	logrus.WithField("deletedCount", deleted).Info("deleted samples by source")
	return nil
}
