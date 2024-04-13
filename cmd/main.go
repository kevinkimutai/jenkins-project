package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kevinkimutai/metadata/internal/adapter/db"
	handler "github.com/kevinkimutai/metadata/internal/adapter/handlers"
	"github.com/kevinkimutai/metadata/internal/adapter/server"
	application "github.com/kevinkimutai/metadata/internal/app/core/api"
)

func main() {
	//Get env var in development
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	// Get database connection details from environment variables
	POSTGRES_USERNAME := os.Getenv("POSTGRES_USERNAME")
	POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	PORT := os.Getenv("APPLICATION_PORT")
	DATABASE_PORT := os.Getenv("DATABASE_PORT")

	//Concatinate DB String
	DBURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		POSTGRES_USERNAME,
		POSTGRES_PASSWORD,
		"localhost",
		DATABASE_PORT,
		"moviedb")

	//Connect To DB
	dbAdapter := db.NewDB(DBURL)

	//Dependency Injection
	movieRepo := application.NewMovieRepo(dbAdapter)
	ratingRepo := application.NewRatingRepo(dbAdapter)

	//Services
	ratingService := handler.NewRatingService(ratingRepo)
	movieService := handler.NewMovieService(movieRepo)

	//Server
	server := server.New(PORT, movieService, ratingService)
	server.Run()
}
