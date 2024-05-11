package indexer

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/fs"
	"sierra/common/sha256"
	"sierra/common/uri"
	"sierra/services/sierra/internal/analyzers/format"
	"sierra/services/sierra/internal/modules/sample"
	"sierra/services/sierra/internal/modules/source"
	"sierra/services/sierra/internal/modules/sourcesample"
	"sierra/services/sierra/internal/sierradb"
)

type Indexer struct {
}

func New() *Indexer {
	return &Indexer{}
}

func (indexer *Indexer) Index(ctx context.Context, source source.Source, forceReindex bool) error {
	err := source.Walk(func(fileUri *uri.URI, info fs.FileInfo, err error) error {
		if err != nil {
			logrus.WithError(err).Error("error walking source")
			return nil
		}

		if info.IsDir() {
			return nil
		}

		err = indexFile(ctx, source, fileUri, forceReindex)
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

func indexFile(ctx context.Context, src source.Source, fileUri *uri.URI, forceReindex bool) error {
	file, err := src.Open(fileUri.Path())
	if err != nil {
		return fmt.Errorf("failed opening file: %w", err)
	}
	defer file.Close()

	sha256Str, err := sha256.HashHexEncoded(file)
	if err != nil {
		return fmt.Errorf("failed hashing file: %w", err)
	}

	if !forceReindex {
		// TODO: Don't fetch one by one
		_, err = sample.Get(ctx, *sha256Str)
		if err == nil {
			// TODO: Make a smarter caching mechanism
			// Don't reindex file
			return nil
		}
		if !sierradb.IsNotFound(err) {
			return fmt.Errorf("failed getting sample: %w", err)
		}
	}

	// Format
	foundFormat, found, err := format.AnalyzeFormat(file)
	if err != nil {
		return fmt.Errorf("failed analyzing format: %w", err)
	}

	if !found {
		// Skipping non-supported format
		return nil
	}

	err = sample.Upsert(ctx, *sha256Str, foundFormat.Type)
	if err != nil {
		return fmt.Errorf("failed saving sample to db: %w", err)
	}

	err = sourcesample.Upsert(ctx, *sha256Str, src.GetURI(), fileUri)
	if err != nil {
		return fmt.Errorf("failed saving source sample to db: %w", err)
	}

	return nil
}
