package database

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

func RunMigrations(connString string, db *DB) error {

	driver, err := mysql.WithInstance(db.conn, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://interval/database/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	return nil
}
