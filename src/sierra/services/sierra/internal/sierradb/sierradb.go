package sierradb

import (
	"context"
	"embed"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"sierra/common/database"
	"sierra/services/sierra/config"
	"sierra/services/sierra/internal/sierradb/sierraent"
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

	dbFilePath, err := getSierraDbFilePath()
	if err != nil {
		return nil, fmt.Errorf("failed loading sierra db: %w", err)
	}

	err = createFileIfMissing(dbFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed creating sierra db file: %w", err)
	}

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

func getSierraDbFilePath() (string, error) {
	userCacheDir, err := config.GetAppDataDir()
	if err != nil {
		return "", fmt.Errorf("failed getting user config dir: %w", err)
	}

	return filepath.Join(userCacheDir, sierraDbFileName), nil
}

func createFileIfMissing(dbFilePath string) error {
	err := config.CreateAppDataDir()
	if err != nil {
		return fmt.Errorf("failed creating sierra db dir: %w", err)
	}

	file, err := os.OpenFile(dbFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("failed opening sierra db file: %w", err)
	}
	_ = file.Close()

	return nil
}

func IsNotFound(err error) bool {
	return sierraent.IsNotFound(err)
}
