//go:build migrate

package app

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // драйвер для PostgreSQL
	"github.com/pressly/goose/v3"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

func init() {
	databaseURL, ok := os.LookupEnv("PG_URL")
	if !ok || len(databaseURL) == 0 {
		log.Fatalf("migrate: environment variable not declared: PG_URL")
	}

	databaseURL += "?sslmode=disable"

	// Подключение к базе данных
	var db *sqlx.DB
	var err error

	for attempts := _defaultAttempts; attempts > 0; attempts-- {
		db, err = sqlx.Open("postgres", databaseURL)
		if err == nil {
			break
		}

		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
	}

	if err != nil {
		log.Fatalf("Migrate: postgres connect error: %s", err)
	}

	// Инициализация Goose для выполнения миграций
	err = goose.Up(db.DB, "./migrations") // путь к директории с SQL миграциями
	if err != nil {
		log.Fatalf("Migrate: error applying migrations: %s", err)
	}

	log.Println("Migrations applied successfully")
}
