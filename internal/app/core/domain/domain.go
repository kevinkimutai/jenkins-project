package domain

import (
	"errors"
	"time"
)

type DataResponse struct {
	StatusCode uint        `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
}

type Response struct {
	StatusCode    uint          `json:"status_code"`
	Message       string        `json:"message"`
	Page          uint          `json:"page"`
	NumberOfPages uint          `json:"number_of_pages"`
	Total         uint          `json:"total"`
	Data          []interface{} `json:"data"`
}

type FetchData struct {
	Page          uint          `json:"page"`
	NumberOfPages uint          `json:"number_of_pages"`
	Total         uint          `json:"total"`
	Data          []interface{} `json:"data"`
}

type MovieParams struct {
	MovieStartAirDate string
	MovieEndAirDate   string
	SearchString      string
	Page              string
	Limit             string
}

type Movie struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Director      string    `json:"director"`
	CreatedAt     time.Time `json:"created_at"`
	AverageRating float64   `json:"average_rating"`
	ReleaseDate   time.Time `json:"release_date"`
}

type Rating struct {
	ID        int64     `json:"id"`
	MovieID   int64     `json:"movie_id"`
	Rating    float64   `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}

func NewMovieDomain(movie Movie) (Movie, error) {
	if movie.Title == "" {
		return movie, errors.New("missing title field")
	}
	if movie.Description == "" {
		return movie, errors.New("missing description field")
	}
	if movie.Director == "" {
		return movie, errors.New("missing director field")
	}

	return movie, nil
}

func NewRatingsDomain(rating Rating) (Rating, error) {
	if rating.MovieID == 0 {
		return rating, errors.New("missing movie_id field")
	}
	if rating.Rating == 0 {
		return rating, errors.New("missing rating field")
	}

	return rating, nil
}

func CheckMovieParams(m map[string]string) MovieParams {

	movieParams := MovieParams{}

	if m["movie_start_date"] != "" && m["movie_end_date"] != "" {
		movieParams.MovieStartAirDate = m["movie_start_date"]
		movieParams.MovieEndAirDate = m["movie_end_date"]
	}

	if m["search_string"] != "" {
		movieParams.SearchString = m["search_string"]
	}

	if m["page"] != "" {
		movieParams.Page = m["page"]
	}

	if m["limit"] != "" {
		movieParams.Limit = m["limit"]
	}

	if m["sort"] != "" {
		movieParams.Page = m["sort"]
	}
	return movieParams

}
