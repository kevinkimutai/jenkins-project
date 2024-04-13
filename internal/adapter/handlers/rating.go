package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/metadata/internal/app/core/domain"
	"github.com/kevinkimutai/metadata/internal/ports"
)

type RatingService struct {
	api ports.RatingApiPort
}

func NewRatingService(api ports.RatingApiPort) *RatingService {
	return &RatingService{api: api}
}

// CreateMovieRating registers a new movie rating
// @Summary Create a new movie rating
// @Description Create Rating
// @Tags rating
// @Accept json
// @Produce json
// @Param body domain.Rating true "Create Rating"
// @Success 201 {object} domain.DataResponse "success"
// @Failure 400 {object} domain.ErrorResponse{}
// @Router /rating [post]
func (s *RatingService) CreateMovieRating(c *fiber.Ctx) error {
	rating := domain.Rating{}

	//Bind To struct
	if err := c.BodyParser(&rating); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Check Missing Inputs
	rating, err := domain.NewRatingsDomain(rating)
	if err != nil {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 400,
				Message:    err.Error(),
			})
	}

	//api
	rating, err = s.api.CreateNewRating(rating)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Map Response
	res := domain.DataResponse{
		StatusCode: 201,
		Message:    "success",
		Data:       rating,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
