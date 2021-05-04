package pg

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/phungvandat/source-template/config/env"
	"github.com/phungvandat/source-template/model/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// CreateTestDatabase create database connection to testing
func CreateTestDatabase(t *testing.T) (*gorm.DB, func()) {
	connStr := "host=localhost port=5433 user=postgres password=pass_demo_test dbname=db_demo_test sslmode=disable"
	if env.PGTestingSource() != "" {
		connStr = env.PGTestingSource()
	}

	db, dbErr := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if dbErr != nil {
		t.Fatalf("Fail to create database. %s", dbErr.Error())
	}

	rand.Seed(time.Now().UnixNano())

	schemaName := "test" + strconv.FormatInt(rand.Int63(), 10)

	err := db.Exec("CREATE SCHEMA " + schemaName).Error
	if err != nil {
		t.Fatalf("Fail to create schema. %s", err.Error())
	}

	// set schema for current db connection
	err = db.Exec("SET search_path TO " + schemaName).Error
	if err != nil {
		t.Fatalf("Fail to set search_path to created schema. %s", err.Error())
	}

	return db, func() {
		err := db.Exec("DROP SCHEMA " + schemaName + " CASCADE").Error
		if err != nil {
			t.Fatalf("Fail to drop database. %s", err.Error())
		}
	}
}

// MigrateTables migrate db tables
func MigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.Company{},
		&domain.User{},
	)
}
