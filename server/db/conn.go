package db

import (
	"database/sql"
	"flag"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"github.com/apoxy-dev/proximal/core/log"
	sqlc "github.com/apoxy-dev/proximal/server/db/sql"
)

var (
	dbPath = flag.String("sqlite3_path", "sqlite3.db", "SQLite3 database path.")
)

// DB is a wrapper around sql.DB.
type DB struct {
	db *sql.DB
}

// New returns a new DB.
func New() (*DB, error) {
	if _, err := os.Stat(*dbPath); os.IsNotExist(err) {
		log.Infof("creating new database at %s", *dbPath)
		if err := os.MkdirAll(filepath.Dir(*dbPath), 0755); err != nil {
			return nil, err
		}
		if _, err := os.Create(*dbPath); err != nil {
			return nil, err
		}
	}

	db, err := sql.Open("sqlite3", *dbPath)
	if err != nil {
		return nil, err
	}
	return &DB{
		db: db,
	}, nil
}

// Begin starts a transaction.
func (db *DB) Begin() (*sql.Tx, error) { return db.db.Begin() }

// Query returns a new sqlc.Queries instance.
func (db *DB) Queries() *sqlc.Queries {
	return sqlc.New(db.db)
}

// Close closes the database connection.
func (db *DB) Close() {
	db.db.Close()
}
