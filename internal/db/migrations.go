package db

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var migrations embed.FS

func RunMigrations(dbURL string) error {

	// migration source from embed
    src, err := iofs.New(migrations, "migrations")
    if err != nil {
        return fmt.Errorf("error creating source: %w", err)
    }

    // Open a database/sql connection ONLY for migrations
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        return fmt.Errorf("sql.Open failed: %w", err)
    }
    defer db.Close()
    
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        return fmt.Errorf("driver init failed: %w", err)
    }

    m, err := migrate.NewWithInstance("iofs", src, "postgres", driver)
    if err != nil {
        return fmt.Errorf("migrate init failed: %w", err)
    }

    err = m.Up()
    if err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("migration failed: %w", err)
    }

    log.Println("DB migrations applied")
    return nil
}
