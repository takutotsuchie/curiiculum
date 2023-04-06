package db

import (
	"database/sql"
	"sync"
)

var (
	DB      *sql.DB
	DBMutex sync.Mutex
)

func SetDB(db *sql.DB) {
	DBMutex.Lock()
	defer DBMutex.Unlock()

	DB = db
}

func GetDB() *sql.DB {
	DBMutex.Lock()
	defer DBMutex.Unlock()

	return DB
}
