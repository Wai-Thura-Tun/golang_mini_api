package database

import (
	"context"
	"database/sql"
	"time"
)

type DB struct {
	conn *sql.DB
}

func NewDB(ctx context.Context, connString string) (*DB, error) {
	conn, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	// Test the connection
	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()

	if err = conn.PingContext(ctx); err != nil {
		conn.Close()
		return nil, err
	}

	// Configure the datbase connection pool
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(5)

	return &DB{conn: conn}, nil
}

// Close the database connection
func (db *DB) Close() error {
	return db.conn.Close()
}
