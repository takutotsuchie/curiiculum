package db

import (
	"database/sql"
	"sync"
)

var (
	DB      *sql.DB
	DBMutex sync.Mutex
)

// DBインスタンスをセット
func SetDB(db *sql.DB) {
	DBMutex.Lock()
	defer DBMutex.Unlock()

	DB = db
}

// DBインスタンスを取得
func GetDB() *sql.DB {
	DBMutex.Lock()
	defer DBMutex.Unlock()

	return DB
}
