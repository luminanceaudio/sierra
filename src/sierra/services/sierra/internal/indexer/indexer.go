package indexer

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/fs"
	"os"
	"path/filepath"
	"sierra/common/format"
	"sierra/common/pathfinder"
	"sierra/common/safemap"
	"sierra/common/sha256"
	"sierra/common/uri"
	"sierra/common/waveform"
	"sierra/services/sierra/config"
	"sierra/services/sierra/internal/modules/sample"
	"sierra/services/sierra/internal/modules/source"
	"sierra/services/sierra/internal/modules/sourcesample"
	"sierra/services/sierra/internal/sierradb"
)

var singleton *Indexer

type Indexer struct {
	sourceURIToIndexer *safemap.SafeMap[string, sourceIndexer]
}

func Singleton() *Indexer {
	if singleton != nil {
		return singleton
	}

	singleton = &Indexer{
		sourceURIToIndexer: safemap.New[string, sourceIndexer](),
	}
	return singleton
}

func (i *Indexer) Index(ctx context.Context, source source.Source, forceReindex, inBackground bool) error {
	sourceUri := source.GetURI().String()

	tempSourceIdx, ok := i.sourceURIToIndexer.Get(sourceUri)
	if !ok {
		tempSourceIdx = newSourceIndexer()
		i.sourceURIToIndexer.Put(sourceUri, tempSourceIdx)
	}

	// Don't index if in progress
	if tempSourceIdx.IsIndexing {
		logrus.WithField("sourceUri", sourceUri).Info("not indexing source as it is already being indexed")
		return nil
	}

	tempSourceIdx.IsIndexing = true

	// Save IsIndexing state
	i.sourceURIToIndexer.Put(sourceUri, tempSourceIdx)

	// Run index
	if inBackground {
		go func() {
			err := i.index(ctx, source, forceReindex)
			if err != nil {
				logrus.WithError(err).WithField("sourceUri", sourceUri).Info("failed to index source")
				return
			}

			logrus.WithField("sourceUri", sourceUri).Info("indexed source in background successfully")
		}()
	} else {
		err := i.index(ctx, source, forceReindex)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *Indexer) index(ctx context.Context, source source.Source, forceReindex bool) error {
	defer i.sourceURIToIndexer.Delete(source.GetURI().String())

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
		_, err = sourcesample.GetBySampleSha256(ctx, *sha256Str)
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

	imageCacheDir, err := config.CreateAppImageCacheDir()
	if err != nil {
		return fmt.Errorf("failed creating cache dir: %w", err)
	}

	waveformFilePath := filepath.Join(imageCacheDir, pathfinder.RandomFilename()) + ".svg"
	f, err := os.OpenFile(waveformFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("failed opening file: %w", err)
	}
	defer f.Close()

	err = waveform.DrawWaveformSVG(foundFormat.Samples, f)
	if err != nil {
		logrus.WithError(err).Error("failed drawing waveform")
	}

	err = sample.Upsert(ctx, *sha256Str, foundFormat.Type, waveformFilePath)
	if err != nil {
		return fmt.Errorf("failed saving sample to db: %w", err)
	}

	err = sourcesample.Upsert(ctx, *sha256Str, src.GetURI(), fileUri)
	if err != nil {
		return fmt.Errorf("failed saving source sample to db: %w", err)
	}

	return nil
}
