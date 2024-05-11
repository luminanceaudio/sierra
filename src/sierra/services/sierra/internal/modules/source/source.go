package source

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Source interface {
	Open(name string) (*os.File, error)
	Walk(fn filepath.WalkFunc) error
	WriteFile(filename string, data []byte, perm os.FileMode) error
	GetURI() string
}

func New(uri string) (Source, error) {
	split := strings.SplitN(uri, "://", 2)
	if len(split) != 2 {
		return nil, fmt.Errorf("invalid source uri: %s", uri)
	}

	schema, path := split[0], split[1]

	switch schema {
	case File:
		return NewLocalSource(path)
	}

	return nil, fmt.Errorf("unsupported source type %s", schema)
}
