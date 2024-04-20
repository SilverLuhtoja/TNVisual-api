package test_utils

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/SilverLuhtoja/TNVisual/internal/api"
	"github.com/SilverLuhtoja/TNVisual/internal/database"
)

const TEST_DATABASE_URL = "postgres://postgres:postgres@localhost:15432/postgres?sslmode=disable"

func CreateTestConfig() *api.ApiConfig {
	db := openDatabase()
	return &api.ApiConfig{DB: database.New(db), Client: &http.Client{}}
}

func ClearTable(table string) {
	db := openDatabase()
	statment := fmt.Sprintf("DELETE FROM %s", table)
	_, err := db.Exec(statment)
	if err != nil {
		log.Fatal("Couldn't not delete data from table - ", err)
	}
}

func openDatabase() *sql.DB {
	db, err := sql.Open("postgres", TEST_DATABASE_URL)

	if err != nil {
		log.Fatalf("Couldn't connect with database: %s", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Couldn't ping database: %s", err)
	}

	return db
}
