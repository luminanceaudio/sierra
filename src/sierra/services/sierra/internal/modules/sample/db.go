package sample

import (
	"context"
	"fmt"
	"github.com/luminanceaudio/sierra/src/sierra/common/sha256"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/format"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent/sample"
	"strings"
)

func Get(ctx context.Context, sha256Str string) (*sierraent.Sample, error) {
	if _, err := sha256.HashHexEncoded(strings.NewReader(sha256Str)); err != nil {
		return nil, err
	}

	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return nil, err
	}

	return sierraDb.Client.Sample.Query().Where(sample.ID(sha256Str)).Only(ctx)
}

func Upsert(ctx context.Context, sha256Str string, fileFormat string) error {
	if _, err := sha256.New(sha256Str); err != nil {
		return err
	}

	if !format.IsSupported(fileFormat) {
		return fmt.Errorf("file format %s is not supported", fileFormat)
	}

	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return err
	}

	err = sierraDb.Client.Sample.Create().
		SetID(sha256Str).
		SetFormat(fileFormat).
		OnConflictColumns(sample.FieldID).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("could not upsert sample: %w", err)
	}

	return nil
}
