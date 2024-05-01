package infrastructure

import (
	"database/sql"
	"log"
	"os"

	"github.com/SilverLuhtoja/TNVisual/internal/database"
)

func NewDatabase() *database.Queries {
	db := ConnectToDatabase()
	return database.New(db)
}

func ConnectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Couldn't connect with database: %s", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Couldn't ping database: %s", err)
	}

	return db
}

// func dnsUrl() string {
// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASS")
// 	host := os.Getenv("DB_HOST")
// 	port := os.Getenv("DB_PORT")
// 	databaseName := os.Getenv("DB_NAME")
// 	sslMode := os.Getenv("DB_SSLMODE")

// 	// "postgres://postgres:postgres@localhost:15432/postgres?sslmode=disable"
// 	// postgres://{user}:{password}@{hostname}:{port}/{database-name}?{options}
// 	// return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, databaseName, sslMode)
// 	// return "postgres://postgres:postgres@localhost:15432/postgres?sslmode=disable"
// 	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, databaseName, port, sslMode)
// }
