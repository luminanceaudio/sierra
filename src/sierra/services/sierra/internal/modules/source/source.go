package source

import (
	"fmt"
	"github.com/luminanceaudio/sierra/src/sierra/common/uri"
	"io/fs"
	"os"
)

type WalkFunc func(fileUri *uri.URI, info fs.FileInfo, err error) error

type Source interface {
	Open(name string) (*os.File, error)
	Walk(fn WalkFunc) error
	WriteFile(filename string, data []byte, perm os.FileMode) error
	GetURI() *uri.URI
}

func New(uriStr string) (Source, error) {
	uri, err := uri.New(uriStr)
	if err != nil {
		return nil, err
	}

	switch uri.Scheme() {
	case File:
		return NewLocalSource(uri)
	}

	return nil, fmt.Errorf("unsupported source type %s", uri.Scheme())
}
