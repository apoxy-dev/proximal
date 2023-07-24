package db

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/mattn/go-sqlite3"

	"github.com/apoxy-dev/proximal/server/db/sql/migrations"
)

// DoMigrations performs DB migrations using github.com/golang-migrate/migrate/v4 package.
func DoMigrations() error {
	d, err := iofs.New(migrations.SQL, ".")
	if err != nil {
		return err
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, fmt.Sprintf("sqlite3://%s", *dbPath))
	if err != nil {
		return err
	}
	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("Migrations already up-to-date.")
			return nil
		}
		return err
	}
	log.Println("Migrations complete.")
	return nil
}
