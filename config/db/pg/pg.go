package pg

import (
	"sync"
	"time"

	"github.com/phungvandat/source-template/utils/logger"
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

// InitPGConnection new database connection
func InitPGConn(connStr string) {
	once.Do(func() {
		var err error
		gormDB, err = gorm.Open(postgres.Open(connStr), nil)
		if err != nil {
			panic(err)
		}

		db, err := gormDB.DB()
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		db.SetMaxOpenConns(maxOpenConns)
		db.SetMaxIdleConns(maxIdleConns)
		db.SetConnMaxLifetime(maxConnMaxLifetime)

		logger.Info("pg db connected")
	})
}

// GetDB return a db connection
func GetDB() *gorm.DB {
	if gormDB == nil {
		panic("missed InitPGConnection")
	}
	return gormDB
}

func Close() {
	if gormDB == nil {
		panic("missed InitPGConnection")
	}

	db, err := gormDB.DB()
	if err != nil {
		logger.Error("failed get DB by error: %v", err)
		return
	}

	err = db.Close()
	if err != nil {
		logger.Error("failed close db by error: %v", err)
	}

	logger.Info("pg db closed")
}
