package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	var err error
	if os.Getenv("ENV") != "production" {
		err = godotenv.Load(".env.migrator")
		if err != nil {
			panic(fmt.Sprintf("Failed to load .env file by error: %v", err))
		}
	}

	db, err := sql.Open("postgres", os.Getenv("PG_SOURCE"))
	if err != nil {
		panic(err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var (
		sourceURL  = fmt.Sprintf("%v/migrations", pwd)
		migrations = &migrate.FileMigrationSource{
			Dir: sourceURL,
		}
	)

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
