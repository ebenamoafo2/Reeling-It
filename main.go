package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"reelingit.com/data"
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
	// Initialize the logger
	logInstance := initializeLogger()

	//Environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file was available")
	}

	//Connect to database
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("DATABASE_URL is not set in the environment")
	}
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	//Initialize data repository for the movies
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize movie repository: %v", err)
	}

	//Movie handler initializer
	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)
	http.HandleFunc("/api/movies/search", movieHandler.SearchMovies)
	http.HandleFunc("/api/movies/", movieHandler.GetMovie) // This will handle requests like /api/movies/{id}
	http.HandleFunc("/api/genre/", movieHandler.GetGenres)

	//handle for static files
	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8080"

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)    // Log the error to the console
		logInstance.Error("Failed to start server", err) // Log the error to the file
	}
}
