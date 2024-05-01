package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SilverLuhtoja/TNVisual/internal/route"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // this allows to connect to postgres database
)

func main() {
	godotenv.Load(".env")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Printf("PORT NOT SET")
		PORT = "8000"
	}

	controllers := route.InitializeAllControllers()
	router := route.NewRouter(controllers)
	server := &http.Server{
		Addr:        ":" + PORT,
		Handler:     router,
		ReadTimeout: 5 * time.Second,
	}

	log.Printf("Server  running on: http://localhost%s/\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
