package tests

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/SilverLuhtoja/TNVisual/internal/database"
	"github.com/SilverLuhtoja/TNVisual/tests/test_utils"
)

const TEST_DATABASE_URL = "postgres://test:test@localhost:25432/test?sslmode=disable"

func GetDatabaseQueries() *database.Queries {
	return database.New(openDatabase())
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
	query_statment := fmt.Sprintf(`INSERT INTO %s ( %s) VALUES  (%s)`, tableName, strings.Join(columnNames, ","), test_utils.ValuesToSql(values))

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
