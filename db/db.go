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

	DB = db
}

// DBインスタンスを取得
func GetDB() *sql.DB {
	return DB
}
