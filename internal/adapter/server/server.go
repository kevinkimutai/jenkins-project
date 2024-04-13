package server

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/kevinkimutai/metadata/docs/swagger"
	"github.com/kevinkimutai/metadata/internal/ports"
)

//	@title			Movie API
//	@version		1.0
//	@description	Get movies and its ratings API.
//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8000
//	@BasePath	/api/v1

type ServerAdapter struct {
	port   string
	movie  ports.MovieHandlerPort
	rating ports.RatingHandlerPort
}

func New(port string, movie ports.MovieHandlerPort, rating ports.RatingHandlerPort) *ServerAdapter {
	return &ServerAdapter{port: port, movie: movie, rating: rating}
}

func (s *ServerAdapter) Run() {
	//Initialize Fiber
	app := fiber.New()

	//Logger Middleware
	app.Use(logger.New())

	//Swagger Middleware
	cfg := swagger.Config{
		BasePath: "/api/v1/",
		FilePath: "./docs/swagger/swagger.json",
		Path:     "swagger",
		Title:    "Swagger Movie API Docs",
	}
	app.Use(swagger.New(cfg))

	// Define routes
	app.Route("/api/v1/rating", s.RatingsRouter)
	app.Route("/api/v1/metadata", s.MetadataRouter)

	app.Listen(":" + s.port)
}
