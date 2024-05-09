package sierradb

import (
	"fmt"
	"io/fs"
	"os"
	"sierra/services/sierra/internal/modules/source"
)

const (
	sierraDbFileName = "sierra.db"
)

type SierraDb struct {
	fs.File
}

func New(source source.Source) (*SierraDb, error) {
	sierraDbFile, err := source.Open(sierraDbFileName)
	if err != nil {
		if os.IsNotExist(err) {

		}

		return nil, fmt.Errorf("error opening index file: %w", err)
	}
	defer sierraDbFile.Close()

	return &SierraDb{}
}
