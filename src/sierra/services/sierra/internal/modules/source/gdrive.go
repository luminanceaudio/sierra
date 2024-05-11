package source

import (
	"os"
	"path/filepath"
)

// GDriveSource is a Google Drive source
type GDriveSource struct {
	BaseSource
}

func (s *GDriveSource) WriteFile(filename string, data []byte, perm os.FileMode) error {
	//TODO implement me
	panic("implement me")
}

func (s *GDriveSource) Open(name string) (*os.File, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GDriveSource) Walk(fn filepath.WalkFunc) error {
	//TODO implement me
	panic("implement me")
}

func NewGDriveSource(path string) (*GDriveSource, error) {
	return &GDriveSource{
		BaseSource: newBaseSource(path),
	}, nil
}
