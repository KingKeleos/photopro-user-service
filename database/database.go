package database

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var fs embed.FS

// Migrate database with migration files
func Migrate(PGclient *sql.DB, databaseName string) error {
	m, err := newMigrator(PGclient, databaseName)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		slog.Error("migration", "error", err)
		return err
	}
	if errors.Is(err, migrate.ErrNoChange) {
		slog.Info("migration", "info", err)
	}
	return nil
}

func newMigrator(PGclient *sql.DB, databaseName string) (*migrate.Migrate, error) {
	//new postgreSQL Connection
	driver, err := postgres.WithInstance(PGclient, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("create pg driver: %w", err)
	}
	d, err := iofs.New(fs, "migrations")
	if err != nil {
		return nil, fmt.Errorf("read migrations: %w", err)
	}
	m, err := migrate.NewWithInstance("iofs", d, databaseName, driver)
	if err != nil {
		return nil, fmt.Errorf("create migrate instance: %w", err)
	}

	return m, nil
}
