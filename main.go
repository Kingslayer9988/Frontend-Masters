package main

import (
	"fmt"
	"log"
	"net/http"

	"frontendmasters.com/joho/godotenv"
	"frontendmasters.com/movies/handlers"
	"frontendmasters.com/movies/logger"
)

func initializeLogger() *logger.Logger {
	// Log to File with NewLogger
	logInstance, err := logger.NewLogger("movie-service.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()
	return logInstance
}

func main() {
	// Initialize logger
	logInstance := initializeLogger()
	// Environmental Variables
	if err := godotenv.Load(); err != nil {
		logInstance.Error("Error loading .env file", err)
	}
	godotenv.load()

	//Movie Handler Initialization
	movieHandler := handlers.MovieHandler{}
	// Set up Routes
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

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
