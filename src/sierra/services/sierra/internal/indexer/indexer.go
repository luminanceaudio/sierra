package indexer

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/fs"
	"sierra/services/sierra/internal/analyzers/format"
	"sierra/services/sierra/internal/sierradb/modules/samplefile"
	"sierra/services/sierra/internal/source"
)

type Indexer struct {
}

func New() *Indexer {
	return &Indexer{}
}

func (indexer *Indexer) Index(ctx context.Context, source source.Source) error {
	err := source.Walk(func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			logrus.WithError(err).Error("error walking source")
			return nil
		}

		if info.IsDir() {
			return nil
		}

		err = indexFile(ctx, source, path)
		if err != nil {
			logrus.WithError(err).Error("error indexing file")
			return nil
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed reading directory: %s", err)
	}

	return nil
}

func indexFile(ctx context.Context, source source.Source, filePath string) error {
	file, err := source.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed opening file: %w", err)
	}
	defer file.Close()

	hasher := sha256.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		return fmt.Errorf("failed hashing file: %w", err)
	}
	sha256Str := hex.EncodeToString(hasher.Sum(nil))

	// Format
	foundFormat, found, err := format.AnalyzeFormat(file)
	if err != nil {
		return fmt.Errorf("failed analyzing format: %w", err)
	}

	if !found {
		// Skipping non-supported format
		return nil
	}

	err = samplefile.Create(ctx, sha256Str, foundFormat.Type)
	if err != nil {
		return fmt.Errorf("failed saving sample to db: %w", err)
	}

	return nil
}