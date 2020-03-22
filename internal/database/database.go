package database

import (
	"Instracker/internal/config"
	"github.com/jackc/pgx"
)

// Database is a structure for connection managing
type Database struct {
	Conn *pgx.ConnPool
}

// NewDatabase returns an instance of Database
func NewDatabase(cfg *config.Config) (*Database, error) {
	pgxConfig, err := pgx.ParseURI(cfg.Database.URI)
	if err != nil {
		return nil, err
	}

	db := new(Database)
	if db.Conn, err = pgx.NewConnPool(
		pgx.ConnPoolConfig{
			ConnConfig: pgxConfig,
		}); err != nil {
		return nil, err
	}

	return db, nil
}

// Disconnect closes connection
func (db *Database) Disconnect() {
	db.Conn.Close()
}
