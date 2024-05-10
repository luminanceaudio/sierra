package database

import (
	"embed"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func ApplyMigrations(migrations embed.FS, connectionString string) error {
	migrationsSource, err := iofs.New(migrations, "migrations")
	if err != nil {
		return errors.Wrapf(err, "failed creating new iofs for migrations dir")
	}

	m, err := migrate.NewWithSourceInstance("iofs", migrationsSource, connectionString)
	if err != nil {
		return errors.Wrapf(err, "failed creating new migrations source")
	}
	defer func() {
		_, err = m.Close()
		if err != nil {
			logrus.WithError(err).Error("failed closing migrations source")
		}
	}()

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrapf(err, "failed applying migrations. if error says fix and force, fix your migration, manually run the migration commands that didn't run and set dirty=false in the database migrations table. more info")
	}

	return nil
}
