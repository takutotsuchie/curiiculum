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

// func NewDB() (*sql.DB, error) {

// 	connStr := "host=postgres port=5432 user=postgres password=postgrespassword dbname= postgres sslmode=disable"

// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("DB Open!!!")

// 	return db, nil
// }
