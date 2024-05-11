package sourcesample

import (
	"context"
	"fmt"
	"sierra/services/sierra/internal/sierradb"
	"strings"
)

//
//func GetAll(ctx context.Context) ([]*sierraent.Sample, error) {
//	sierraDb, err := sierradb.Load(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	return sierraDb.Client.Sample.Query().All(ctx)
//}

func Upsert(ctx context.Context, sha256Str, uri string, path string) error {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return err
	}

	relativePath := strings.TrimPrefix(path, uri)

	err = sierraDb.Client.SourceSample.Create().
		SetID(relativePath).
		SetSampleID(sha256Str).
		SetSourceID(uri).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("could not create source sample: %w", err)
	}

	return nil
}
