package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/metadata/internal/app/core/domain"
)

type MovieRepoPort interface {
	GetAllMovies(movieParams domain.MovieParams) (domain.FetchData, error)
	GetMovieById(movieID int64) (*domain.Movie, error)
	CreateMovie(movie domain.Movie) (domain.Movie, error)
}

type MovieApiPort interface {
	GetMovie(movieID int64) (*domain.Movie, error)
	GetMovies(movieParams domain.MovieParams) (domain.FetchData, error)
	CreateNewMovie(movie domain.Movie) (domain.Movie, error)
}

type MovieHandlerPort interface {
	CreateMovie(c *fiber.Ctx) error
	GetAllMovies(c *fiber.Ctx) error
	GetMovieMetadataByID(c *fiber.Ctx) error
}
