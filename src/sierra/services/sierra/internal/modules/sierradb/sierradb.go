package sierradb

import (
	"context"
	"embed"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"sierra/common/database"
	"sierra/services/sierra/internal/modules/sierradb/sierraent"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
)

const (
	sierraUserCacheDirName = "Sierra"
	sierraDbFileName       = "sierra.db"
	sierraDbName           = "sierra"
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

	err = os.MkdirAll(filepath.Dir(dbFilePath), 0700)
	if err != nil {
		return nil, fmt.Errorf("failed to create directory for sierra db %s: %w", filepath.Dir(dbFilePath), err)
	}

	// Create file if doesn't exist
	file, err := os.OpenFile(dbFilePath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed opening sierra db file: %w", err)
	}
	file.Close()

	connectionString := fmt.Sprintf("sqlite3://%s", dbFilePath)

	d, err := iofs.New(migrations, "migrations")
	if err != nil {
		return nil, errors.Wrapf(err, "failed creating new iofs for migrations dir")
	}

	err = database.ApplyMigrations(ctx, d, dialect.SQLite, sierraDbName, connectionString)
	if err != nil {
		return nil, errors.Wrapf(err, "failed applying migrations")
	}

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
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("failed getting user config dir: %w", err)
	}

	return filepath.Join(userCacheDir, sierraUserCacheDirName, sierraDbFileName), nil
}
