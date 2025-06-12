package main

import (
	"fmt"
	"log"
	"net/http"

	"frontendmasters.com/movies/handlers"
	"frontendmasters.com/movies/logger"
)

func main() {
	// Initialize logger
	logInstance := initializeLogger()

	movieHandler := handlers.MovieHandler{}
	// Set up Routes
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)

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

// Call other package ../logger/
func initializeLogger() *logger.Logger {
	// Log to File with NewLogger
	logInstance, err := logger.NewLogger("movie-service.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()
	return logInstance
}
