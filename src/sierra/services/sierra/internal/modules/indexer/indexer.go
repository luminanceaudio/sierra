package indexer

import (
	"sierra/services/sierra/internal/modules/sierradb"
	"sierra/services/sierra/internal/modules/source"
)

type Indexer struct {
	source source.Source
}

func New(source source.Source) (*Indexer, error) {
	sierraDb, err := sierradb.New(source)
	if err != nil {
		
	}
	defer sierraDbFile.Close()

	//indexFile.

	return &Indexer{
		source: source,
	}, nil
}

func (indexer *Indexer) Index(message string) {

}
