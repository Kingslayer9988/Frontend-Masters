package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"frontendmasters.com/movies/data"
	"frontendmasters.com/movies/handlers"
	"frontendmasters.com/movies/logger"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie.log")
	// logInstance.Error("Hello from the Error system", nil)
	if err != nil {
		log.Fatalf("Failed to initialice logger %v", err)
	}
	return logInstance
}

func main() {
	// Initialize logger
	logInstance := initializeLogger()

	// Environmental Variables
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or failed to load: %v", err)
	}

	// Database connection
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatalf("DATABASE_URL not set in environment")
	}
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Data Repository for Movies
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize movierepository")
	}

	// Initialize Movie Repository
	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)

	http.HandleFunc("/api/movies/top/", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random/", movieHandler.GetRandomMovies)
	http.HandleFunc("/api/movies/search/", movieHandler.SearchMovies)
	http.HandleFunc("/api/movies/", movieHandler.GetMovie) // api/movies/140
	http.HandleFunc("/api/genres/", movieHandler.GetGenres)

	// Handler for static files (Frontend)
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Serving the files")

	// Start server
	const addr = ":4020"
	// Log to Console
	logInstance.Info("Server starting on " + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server failed to start", err)
		log.Fatalf("Server failed: %v", err)
	}
}
