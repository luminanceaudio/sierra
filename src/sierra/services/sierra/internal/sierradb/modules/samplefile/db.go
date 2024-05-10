package samplefile

import (
	"context"
	"fmt"
	"sierra/services/sierra/internal/format"
	"sierra/services/sierra/internal/sierradb"
	"sierra/services/sierra/internal/sierradb/internal/sierraent/samplefile"
)

const (
	sha256Length = 64
)

func Create(ctx context.Context, sha256 string, fileFormat string) error {
	if len(sha256) != sha256Length {
		return fmt.Errorf("sha256 length must be %d", sha256Length)
	}

	if !format.IsSupported(fileFormat) {
		return fmt.Errorf("file format %s is not supported", fileFormat)
	}

	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return err
	}

	err = sierraDb.Client.SampleFile.Create().
		SetSha256(sha256).
		SetFormat(fileFormat).
		OnConflictColumns(samplefile.FieldSha256).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("could not create sample file: %w", err)
	}

	return nil
}
