package test_utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/SilverLuhtoja/TNVisual/internal/api"
	"github.com/SilverLuhtoja/TNVisual/internal/database"
)

const TEST_DATABASE_URL = "postgres://test:test@localhost:25432/test?sslmode=disable"

func CreateTestConfig() *api.ApiConfig {
	db := openDatabase()
	return &api.ApiConfig{DB: database.New(db)}
}

// Will help to clean database after each testcase
func ClearTable(table string) {
	db := openDatabase()
	statment := fmt.Sprintf("DELETE FROM %s", table)
	_, err := db.Exec(statment)
	if err != nil {
		log.Fatal("Couldn't not delete data from table - ", err)
	}
}

// Ables to mock data in test tabase
func InsertData(tableName string, columnNames []string, values []string) {
	db := openDatabase()
	query_statment := fmt.Sprintf(`INSERT INTO %s ( %s) VALUES  (%s)`, tableName, strings.Join(columnNames, ","), valuesToSql(values))

	stmt, err := db.Prepare(query_statment)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	stmt.Close()
}

func GetParamsFromResponseBody[T interface{}](structBody T, r *http.Response) T {
	decoder := json.NewDecoder(r.Body)

	params := structBody
	err := decoder.Decode(&params)
	if err != nil {
		log.Fatal(err)
	}

	return params
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

func valuesToSql(values []string) string {
	var vals string
	for _, val := range values {
		vals += fmt.Sprintf("'%s',", val)
	}
	return vals[:len(vals)-1]
}
