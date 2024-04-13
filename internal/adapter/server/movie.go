package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *ServerAdapter) MetadataRouter(api fiber.Router) {
	api.Post("/", s.movie.CreateMovie)
	api.Get("/", s.movie.GetAllMovies)
	api.Get("/:movieID", s.movie.GetMovieMetadataByID)

}
