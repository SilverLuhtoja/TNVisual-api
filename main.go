package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SilverLuhtoja/TNVisual/internal/api"
	"github.com/SilverLuhtoja/TNVisual/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // this allows to connect to postgres database
)

func connectToDatabase() *sql.DB {
	// postgres://{user}:{password}@{hostname}:{port}/{database-name}?{options}
	// connectionPath := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Couldn't connect with database: %s", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Couldn't ping database: %s", err)
	}

	return db
}

func main() {
	godotenv.Load(".env")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Printf("PORT NOT SET")
		PORT = "8000"
	}
	db := connectToDatabase()
	apiConfig := &api.ApiConfig{DB: database.New(db), Client: &http.Client{}}
	router := api.NewRouter(apiConfig)

	server := &http.Server{
		Addr:        ":" + PORT,
		Handler:     router,
		ReadTimeout: 5 * time.Second,
	}

	log.Printf("Server  running on: http://localhost%s/\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
