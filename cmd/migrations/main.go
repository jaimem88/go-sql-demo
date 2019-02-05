// nolint
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	user   = os.Getenv("DB_USER")
	pass   = os.Getenv("DB_PASS")
	host   = os.Getenv("DB_HOST")
	port   = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")
)

func main() {
	// https://github.com/golang-migrate/migrate
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?application_name=%s&sslmode=disable", user, pass, host, port, dbName, "demo")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to sql.Open"))

	}
	defer db.Close()
	driver, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: "demo_migrations",
		DatabaseName:    "postgres",
	})

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrations/schema",
		"postgres", driver)

	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create migrate instance"))
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		if v, ok := err.(migrate.ErrDirty); ok {
			if err := m.Force(v.Version); err != nil {
				log.Fatal(errors.Wrap(err, "failed to force migration - fix your sql code"))
			}
		} else {
			log.Fatal(errors.Wrap(err, "failed to run migrations"))
		}
	}
	log.Println("done.")
}
