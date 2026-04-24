package main

import (
	"log"
	"net/http"

	"reelingit.com/handlers"
	"reelingit.com/logger"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie-service.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logInstance.Close()
	return logInstance
}

func main() {

	logInstance := initializeLogger()
	

	movieHandler := handlers.MovieHandler{}

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

	//handle for static files
	http.Handle("/", http.FileServer(http.Dir("public")))



	const addr = ":8080"

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)    // Log the error to the console
		logInstance.Error("Failed to start server", err) // Log the error to the file
	}
}
