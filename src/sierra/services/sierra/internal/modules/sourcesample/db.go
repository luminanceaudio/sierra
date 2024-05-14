package sourcesample

import (
	"context"
	"fmt"
	"github.com/luminanceaudio/sierra/src/sierra/common/sha256"
	"github.com/luminanceaudio/sierra/src/sierra/common/uri"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/client/models"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent"
	sample2 "github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent/sample"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent/source"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent/sourcesample"
	"github.com/sirupsen/logrus"
	"strings"
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

	return translateSample(sourceSample), nil
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

	return translateSample(sourceSample), nil
}

func getAll(ctx context.Context, query string, page, size *int) (*sierraent.SourceSampleQuery, error) {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return nil, err
	}

	q := sierraDb.Client.SourceSample.Query().
		WithSource().
		WithSample()

	if page != nil && size != nil {
		if *page <= 0 {
			return nil, fmt.Errorf("page must be a positive integer")
		}

		if *size <= 0 {
			return nil, fmt.Errorf("size must be a positive integer")
		}

		q = q.Offset((*page - 1) * *size).Limit(*size)
	}

	if query != "" {
		q = q.Where(sourcesample.RelativePathContains(query))
	}

	return q, nil
}

func Query(ctx context.Context, query string, page, size int) ([]*models.Sample, error) {
	q, err := getAll(ctx, query, &page, &size)
	if err != nil {
		return nil, err
	}

	sourceSamples, err := q.All(ctx)
	if err != nil {
		return nil, err
	}

	return translateSamples(sourceSamples), nil
}

func QueryCount(ctx context.Context, query string) (int, error) {
	q, err := getAll(ctx, query, nil, nil)
	if err != nil {
		return 0, err
	}

	count, err := q.Count(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func Upsert(ctx context.Context, sha256Str string, sourceUri, fileUri *uri.URI) error {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return err
	}

	splitPath := strings.SplitN(fileUri.String(), sourceUri.String(), 2)
	if len(splitPath) != 2 {
		return fmt.Errorf("invalid file uri, can't calculate relative path")
	}
	relativePath := strings.TrimPrefix(splitPath[1], "/")

	err = sierraDb.Client.SourceSample.Create().
		SetID(fileUri.String()).
		SetSampleID(sha256Str).
		SetSourceID(sourceUri.String()).
		SetRelativePath(relativePath).
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
