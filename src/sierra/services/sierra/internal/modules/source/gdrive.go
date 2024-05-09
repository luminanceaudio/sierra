package source

import (
	"io/fs"
)

// GDriveSource is a Google Drive source
type GDriveSource struct {
	BaseSource
}

func (s *GDriveSource) WriteFile(filename string, data []byte, perm fs.FileMode) error {
	//TODO implement me
	panic("implement me")
}

func (s *GDriveSource) Open(name string) (fs.File, error) {
	//TODO implement me
	panic("implement me")
}

func (s *GDriveSource) ReadDir(name string) ([]fs.DirEntry, error) {
	//TODO implement me
	panic("implement me")
}

func NewGDriveSource(path string) (*OneofSource, error) {
	return &OneofSource{
		GDrive: &GDriveSource{
			BaseSource: newBaseSource(path),
		},
	}, nil
}
