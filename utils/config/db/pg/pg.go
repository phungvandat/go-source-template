package pg

import (
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once   = sync.Once{}
	gormDB *gorm.DB
)

const (
	maxOpenConns       = 25
	maxIdleConns       = 25
	maxConnMaxLifetime = time.Minute * 5
)

// InitDBConnection new database connection
func InitDBConnection(connStr string) {
	var err error
	once.Do(func() {
		gormDB, err = gorm.Open(postgres.Open(connStr), nil)
		if err != nil {
			panic(err)
		}

		db, err := gormDB.DB()
		if err != nil {
			panic(err)
		}

		db.SetMaxOpenConns(maxOpenConns)
		db.SetMaxIdleConns(maxIdleConns)
		db.SetConnMaxLifetime(maxConnMaxLifetime)
	})
}

// GetDB return a db connection
func GetDB() *gorm.DB {
	if gormDB == nil {
		panic("missed InitDBConnection")
	}
	return gormDB
}
