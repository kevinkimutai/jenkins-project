package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *ServerAdapter) RatingsRouter(api fiber.Router) {
	api.Post("/", s.rating.CreateMovieRating)
}
