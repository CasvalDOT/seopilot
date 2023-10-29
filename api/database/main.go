package database

import (
	"log"

	_ "github.com/lib/pq"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(connection string) (Database, error) {
	instance, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return Database{}, err
	}

	m, err := migrate.New(
		"file:///database/migrations",
		connection)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
	}

	return Database{
		instance,
	}, err
}
