package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/metadata/internal/app/core/domain"
)

type RatingRepoPort interface {
	CreateRating(rating domain.Rating) (domain.Rating, error)
}

type RatingApiPort interface {
	CreateNewRating(rating domain.Rating) (domain.Rating, error)
}
type RatingHandlerPort interface {
	CreateMovieRating(c *fiber.Ctx) error
}
