package database

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

func ApplyMigrations(ctx context.Context, d source.Driver, driverName, databaseName, connectionString string) error {
	err := createDatabaseIfNotExists(ctx, driverName, databaseName, connectionString)
	if err != nil {
		return errors.Wrapf(err, "failed creating database")
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, connectionString)
	if err != nil {
		return errors.Wrapf(err, "failed creating new migrations source")
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrapf(err, "failed applying migrations. if error says fix and force, fix your migration, manually run the migration commands that didn't run and set dirty=false in the database migrations table. more info")
	}

	return nil
}

func createDatabaseIfNotExists(ctx context.Context, driverName, databaseName, connectionString string) error {
	//db, err := sql.Open(driverName, connectionString)
	//if err != nil {
	//	return errors.Wrapf(err, "failed connecting to postgres")
	//}
	//
	//_, err = db.ExecContext(ctx, fmt.Sprintf(`CREATE DATABASE %s`, databaseName))
	//if err != nil {
	//	// TODO: Figure out how to not fail in sqlite3 when db exists
	//	//if err, ok := err.(*sqlite3.Error); ok && err.Code.Error() != sqlite3.err {
	//	return errors.Wrapf(err, "failed creating database '%s'", databaseName)
	//	//}
	//}

	return nil
}
