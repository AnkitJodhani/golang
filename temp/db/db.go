package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB // Database connection DB.

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name or DNS or ENDPOINT", err)
		panic(err)
	}
	// defer DB.Close() // ⚠️ DON'T DO SUCH MISTAKE - it will close the database
	DB.SetConnMaxLifetime(0)
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)
	createTable()
}

func createTable() {
	createEventTableQuery := `
		CREATE TABLE IF NOT EXISTS events (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL
		)`

	result, err := DB.Exec(createEventTableQuery)
	if err != nil {
		log.Fatal("there was an error while creating table", err)
		panic(err)
	}
	fmt.Println("tabel create successfully", result)

}
