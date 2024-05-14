package source

import (
	"fmt"
	"github.com/luminanceaudio/sierra/src/sierra/common/uri"
	"io/fs"
	"os"
	"path/filepath"
)

type LocalSource struct {
	uri *uri.URI
}

func NewLocalSource(uri *uri.URI) (*LocalSource, error) {
	fileStat, err := os.Stat(uri.Path())
	if err != nil {
		return nil, fmt.Errorf("error getting stat of local source: %w", err)
	}

	if !fileStat.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", uri.Path())
	}

	return &LocalSource{
		uri: uri,
	}, nil
}

func (s *LocalSource) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

func (s *LocalSource) Open(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDONLY, 0600)
}

func (s *LocalSource) Walk(fn WalkFunc) error {
	return filepath.Walk(s.uri.Path(), func(path string, info fs.FileInfo, err error) error {
		fileUri := uri.NewFromParts(s.uri.Scheme(), path)
		return fn(fileUri, info, err)
	})
}

func (s *LocalSource) GetURI() *uri.URI {
	return s.uri
}
