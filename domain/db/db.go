// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// SQLDB wraps a standard sql.DB instance and provides
// lifecycle management for SQL database connections.
type SQLDB struct {
	db *sql.DB
}

// NewSQLDB initializes and returns a new SQLDB instance
// using the provided configuration. It sets connection pool parameters
// such as maximum open/idle connections and connection lifetime.
//
// Returns an error if the connection could not be established.
func NewSQLDB(cfg *Config) (*SQLDB, error) {
	db, err := Connect(cfg.Driver, cfg.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(int(cfg.PoolSize))
	db.SetMaxIdleConns(int(cfg.MaxIdle))
	db.SetConnMaxLifetime(cfg.MaxLifetime)

	return &SQLDB{
		db: db,
	}, nil
}

// DB returns a copy of the raw db
func (db *SQLDB) DB() *sql.DB {
	return db.db
}

// Close closes the underlying database connection.
// It is safe to call Close multiple times; if the DB is already nil, it does nothing.
func (db *SQLDB) Close() error {
	if db.db != nil {
		return db.db.Close()
	}
	return nil
}

// Connect opens a new database connection using the given driver name and DSN (Data Source Name).
// It pings the database to verify the connection is valid before returning.
//
// Returns an error if the connection could not be opened or pinged.
func Connect(driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// Ping verifies that the database connection is still alive.
// It is useful for health checks or readiness probes.
func (db *SQLDB) Ping() error {
	return db.db.Ping()
}
