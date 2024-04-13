package handler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/metadata/internal/app/core/domain"
	"github.com/kevinkimutai/metadata/internal/ports"
)

type MovieService struct {
	api ports.MovieApiPort
}

func NewMovieService(api ports.MovieApiPort) *MovieService {
	return &MovieService{api: api}
}

func (s *MovieService) CreateMovie(c *fiber.Ctx) error {
	movie := domain.Movie{}

	//Bind To struct
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Check Missing Inputs
	movie, err := domain.NewMovieDomain(movie)
	if err != nil {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 400,
				Message:    err.Error(),
			})
	}

	//api
	movie, err = s.api.CreateNewMovie(movie)
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
		Data:       movie,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (s *MovieService) GetAllMovies(c *fiber.Ctx) error {

	//Get Query Params
	m := c.Queries()

	//Bind To movieParams
	movieParams := domain.CheckMovieParams(m)

	//api
	data, err := s.api.GetMovies(movieParams)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Map Response
	res := domain.Response{
		StatusCode:    200,
		Message:       "success",
		Page:          data.Page,
		NumberOfPages: data.NumberOfPages,
		Total:         data.Total,
		Data:          data.Data,
	}

	return c.Status(fiber.StatusOK).JSON(res)

}

func (s *MovieService) GetMovieMetadataByID(c *fiber.Ctx) error {
	movieID := c.Params("movieID")

	//Check If MovieID exists
	if movieID == "" {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    "missing movieID",
			})
	}

	//Convert To Int64
	ID, err := strconv.ParseInt(movieID, 10, 64)
	if err != nil {
		log.Fatal("error converting to int64 %w", err)
	}

	movie, err := s.api.GetMovie(ID)

	if err != nil {
		//Dismiss No Movie With ID Error
		if err.Error() == "no rows in result set" {
			res := domain.ErrorResponse{
				StatusCode: 200,
				Message:    fmt.Sprintf("No Movie With ID:%v", movieID),
			}

			return c.Status(500).JSON(res)
		}

		//Map Error
		res := domain.ErrorResponse{
			StatusCode: 500,
			Message:    err.Error(),
		}

		return c.Status(500).JSON(res)
	}

	//Map Response
	res := domain.DataResponse{
		StatusCode: 200,
		Message:    "success",
		Data:       &movie,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
