package sierradb

import (
	"context"
	"embed"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/luminanceaudio/sierra/src/sierra/common/database"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/config"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

const (
	sierraDbFileName = "sierra.db"
)

//go:embed migrations/*.sql
var migrations embed.FS

var sierraDb *SierraDb

type SierraDb struct {
	Client *sierraent.Client
}

func Load(ctx context.Context) (*SierraDb, error) {
	if sierraDb != nil {
		return sierraDb, nil
	}

	dbDirPath, err := config.CreateAppDataDir("")
	if err != nil {
		return nil, fmt.Errorf("failed getting user config dir: %w", err)
	}

	dbFilePath := filepath.Join(dbDirPath, sierraDbFileName)

	file, err := os.OpenFile(filepath.Clean(dbFilePath), os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed creating sierra db file: %w", err)
	}
	_ = file.Close()

	connectionString := fmt.Sprintf("sqlite3://%s", dbFilePath)

	err = database.ApplyMigrations(migrations, connectionString)
	if err != nil {
		return nil, errors.Wrapf(err, "failed applying migrations")
	}

	connectionString = fmt.Sprintf("file:%s", dbFilePath)
	client, err := sierraent.Open(dialect.SQLite, connectionString)
	if err != nil {
		return nil, errors.Wrapf(err, "failed opening connection to postgres")
	}

	sierraDb = &SierraDb{
		Client: client,
	}
	return sierraDb, nil
}

func IsNotFound(err error) bool {
	return sierraent.IsNotFound(err)
}
