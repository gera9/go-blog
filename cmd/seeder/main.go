package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:example@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Open: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{
		DatabaseName: "postgres",
	})
	if err != nil {
		log.Fatalf("WithInstance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("NewWithDatabaseInstance: %v", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Up: %v", err)
	}

	srcErr, dbErr := m.Close()
	if srcErr != nil {
		log.Fatalf("srcErr: %v", srcErr)
	}

	if dbErr != nil {
		log.Fatalf("dbErr: %v", dbErr)
	}

	log.Println("Migrations applied successfully")
}
