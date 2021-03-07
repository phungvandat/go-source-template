package env

import (
	"os"
	"sync"
)

var (
	once         sync.Once
	isProduction bool
)

// IsProduction check production environment
func IsProduction() bool {
	once.Do(func() {
		isProduction = GetENV() == "production"
	})

	return isProduction
}

// GetENV func
func GetENV() string {
	return os.Getenv("ENV")
}
