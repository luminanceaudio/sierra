package collection

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent/collection"
)

func CreateIfNotExist(ctx context.Context, uuid uuid.UUID, name string) error {
	sierraDb, err := sierradb.Load(ctx)
	if err != nil {
		return err
	}

	err = sierraDb.Client.Collection.Create().
		SetID(uuid).
		SetName(name).
		OnConflictColumns(collection.FieldID).
		Ignore().
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("could not upsert collection: %w", err)
	}

	return nil
}
