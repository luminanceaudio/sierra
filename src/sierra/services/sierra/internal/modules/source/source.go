package source

import (
	"encoding/json"
	"fmt"
	"io/fs"
)

type Source interface {
	fs.ReadDirFS
	WriteFile(filename string, data []byte, perm fs.FileMode) error
}

type OneofSource struct {
	Local  *LocalSource  `json:"local"`
	GDrive *GDriveSource `json:"gdrive"`
}

func (s *OneofSource) UnmarshalJSON(bytes []byte) error {
	// Use alias to avoid recursive unmarshaling
	type sourceUnmarshal *OneofSource
	err := json.Unmarshal(bytes, (sourceUnmarshal)(s))
	if err != nil {
		return err
	}

	// If more than one source type is provided, fail
	if s.Local != nil && s.GDrive != nil {
		return fmt.Errorf("cannot specify multiple source types")
	}

	return nil
}

func (s *OneofSource) Get() (Source, error) {
	if s.Local != nil {
		return s.Local, nil
	}

	if s.GDrive != nil {
		return s.GDrive, nil
	}

	return nil, fmt.Errorf("no source provided")
}
