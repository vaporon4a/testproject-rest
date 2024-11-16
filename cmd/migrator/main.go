package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var connString, migrationPath, migrationTable string

	flag.StringVar(&connString, "conn-string", "", "Connection string to DB [postgres://user:password@host:port/dbname?query]")
	flag.StringVar(&migrationPath, "migration-path", "", "path to migrations files")
	flag.StringVar(&migrationTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse()

	if connString == "" {
		panic("conn-string is required")
	}
	if migrationPath == "" {
		panic("migration-path is required")
	}

	m, err := migrate.New(
		"file://"+migrationPath,
		fmt.Sprintf("%s?x-migrations-table=%s&sslmode=disable", connString, migrationTable),
	)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migration required")
			return
		}
		panic(err)
	}
	fmt.Println("all migration applied")
}
