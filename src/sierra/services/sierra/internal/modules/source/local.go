package source

import (
	"fmt"
	"os"
	"path/filepath"
)

type LocalSource struct {
	BaseSource
}

func (s *LocalSource) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

func (s *LocalSource) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func (s *LocalSource) Walk(fn filepath.WalkFunc) error {
	return filepath.Walk(s.Path, fn)
}

func NewLocalSource(path string) (*OneofSource, error) {
	fileStat, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("error getting stat of local source: %w", err)
	}

	if !fileStat.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", path)
	}

	return &OneofSource{
		Local: &LocalSource{
			BaseSource: newBaseSource(path),
		},
	}, nil
}
