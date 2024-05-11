package source

import (
	"github.com/google/uuid"
)

type BaseSource struct {
	Id   uuid.UUID `json:"id"`
	Path string    `json:"path"`
}

func newBaseSource(path string) BaseSource {
	return BaseSource{
		Id:   uuid.New(),
		Path: path,
	}
}

func (s BaseSource) GetURI() string {
	return s.Path
}
