package db

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kevinkimutai/metadata/internal/adapter/queries"
	"github.com/kevinkimutai/metadata/internal/app/core/domain"
)

func (a *DBAdapter) CreateMovie(movie domain.Movie) (domain.Movie, error) {

	//Map to CreateMovieParamsStruct
	movieParams := queries.CreateMovieParams{
		Title:       movie.Title,
		Director:    movie.Director,
		Description: movie.Description,
	}

	mov, err := a.queries.CreateMovie(a.ctx, movieParams)

	//Map structs
	data := domain.Movie{
		ID:          mov.ID,
		Description: mov.Description,
		Title:       mov.Title,
		Director:    mov.Director,
		CreatedAt:   mov.CreatedAt.Time,
	}

	return data, err

}

// Define named query in your sqlc queries file (query.sql)

func (a *DBAdapter) GetAllMovies(movieParams domain.MovieParams) (domain.FetchData, error) {
	args, err := getParams(movieParams)
	if err != nil {
		return domain.FetchData{}, err
	}

	//Movie records based on params
	movies, err := a.queries.ListMovies(a.ctx, args)
	if err != nil {
		return domain.FetchData{}, err
	}

	//Total Count Records
	totalMovies, err := a.queries.CountMovies(a.ctx)
	if err != nil {
		return domain.FetchData{}, err
	}

	var data []interface{}
	for _, movie := range movies {
		data = append(data, movie)
	}

	return domain.FetchData{
		Page:          getPage(args.Offset, args.Limit),
		NumberOfPages: uint(math.Ceil(float64(totalMovies) / float64(args.Limit))),
		Total:         uint(totalMovies),
		Data:          data,
	}, nil

}

func getParams(movieParams domain.MovieParams) (queries.ListMoviesParams, error) {
	var searchStr pgtype.Text
	var release_date_start, release_date_end pgtype.Date
	var limit, offset int32 = 10, 0

	//Get Params
	if movieParams.SearchString != "" {
		searchStr.Scan(movieParams.SearchString)
	}

	if movieParams.MovieStartAirDate != "" {
		//convert to type time
		date, err := time.Parse("2006-01-02", movieParams.MovieStartAirDate)
		release_date_start.Scan(date)
		if err != nil {
			return queries.ListMoviesParams{}, err
		}
	}

	if movieParams.MovieEndAirDate != "" {
		//convert to type time
		date, err := time.Parse("2006-01-02", movieParams.MovieEndAirDate)
		release_date_end.Scan(date)
		if err != nil {
			return queries.ListMoviesParams{}, err
		}
	}

	if movieParams.Limit != "" {
		items, err := strconv.Atoi(movieParams.Limit)

		limit = int32(items)
		if err != nil {
			return queries.ListMoviesParams{}, err
		}
	}
	if movieParams.Page != "" {
		page, err := strconv.Atoi(movieParams.Page)

		if page < 1 {
			page = 1
		}

		offset = (int32(page) - 1) * limit
		if err != nil {
			return queries.ListMoviesParams{}, err
		}
	}

	return queries.ListMoviesParams{
		Column1:       searchStr,
		ReleaseDate:   release_date_start,
		ReleaseDate_2: release_date_end,
		Limit:         limit,
		Offset:        offset,
	}, nil
}

func getPage(offset, limit int32) uint {
	return uint((offset / limit) + 1)
}

func (a *DBAdapter) GetMovieById(movieID int64) (*domain.Movie, error) {
	movie, err := a.queries.GetMovie(a.ctx, movieID)

	avgRating, ok := interfaceToFloat64(movie.AverageRating)
	if ok != nil {
		return &domain.Movie{}, err
	}

	//Map Struct
	data := &domain.Movie{
		ID:            movie.ID,
		Title:         movie.Title,
		Description:   movie.Description,
		Director:      movie.Director,
		AverageRating: avgRating,
		CreatedAt:     movie.CreatedAt.Time,
	}

	if err != nil {
		//Dismiss No Movie With ID Error
		if err.Error() == "no rows in result set" {
			err = nil
			return data, err
		}
		return data, err

	}

	return data, nil

}

func interfaceToFloat64(value interface{}) (float64, error) {
	// Check if the value is already a float64
	if f, ok := value.(float64); ok {
		return f, nil
	}

	// Check if the value is of type pgtype.Numeric
	if numeric, ok := value.(pgtype.Numeric); ok {
		fval, err := numeric.Value()
		if err != nil {
			return 0, err
		}

		//Convert To Float64
		var floatVal float64
		if strVal, ok := fval.(string); ok {
			floatVal, err = strconv.ParseFloat(strVal, 64)
			if err != nil {
				return 0, err
			}
		} else {
			// Handle the case where fval is not a string
			return 0, fmt.Errorf("value is not a string")
		}

		return floatVal, nil
	}

	// If not float64 or pgtype.Numeric, return error
	return 0, errors.New("value cannot be converted to float64")
}
