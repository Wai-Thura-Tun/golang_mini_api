package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	conn *sql.DB
}

func NewDB(ctx context.Context, connString string) (*DB, error) {
	var conn *sql.DB
	var err error

	log.Print(connString)

	for i := 0; i < 5; i++ {
		conn, err = sql.Open("mysql", connString)
		if err != nil {
			log.Printf("Attempt %d: failed to open DB: %v\n", i+1, err)
			time.Sleep(3 * time.Second)
			continue
		}

		pingCtx, cancel := context.WithTimeout(ctx, 4*time.Second)
		err = conn.PingContext(pingCtx)
		cancel()

		if err == nil {
			log.Println("✅ Connected to MySQL")

			// Configure the datbase connection pool
			conn.SetMaxOpenConns(10)
			conn.SetMaxIdleConns(5)

			return &DB{conn: conn}, nil
		}

		log.Printf("Attempt %d: waiting for MySQL... %v\n", i+1, err)
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("could not connect to mysql after retries: %w", err)
}

// Get connection
func (db *DB) Conn() *sql.DB {
	return db.conn
}

// Close the database connection
func (db *DB) Close() error {
	return db.conn.Close()
}

func (db *DB) CheckUserExists(userID uint) (bool, error) {
	row := db.conn.QueryRow("select 1 from users where id = ? limit 1", userID)
	var exists int
	err := row.Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
