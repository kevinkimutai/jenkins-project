package db

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kevinkimutai/metadata/internal/adapter/queries"
	"github.com/kevinkimutai/metadata/internal/app/core/domain"
)

func (a *DBAdapter) CreateRating(rating domain.Rating) (domain.Rating, error) {
	var numeric pgtype.Numeric

	numeric.Scan(rating.Rating)
	//Map Struct
	ratings := queries.CreateRatingParams{
		MovieID: rating.MovieID,
		Rating:  numeric,
	}

	dbRating, err := a.queries.CreateRating(a.ctx, ratings)

	//Map Struct
	data := domain.Rating{
		ID:        dbRating.ID,
		MovieID:   dbRating.MovieID,
		Rating:    rating.Rating,
		CreatedAt: dbRating.CreatedAt.Time,
	}

	return data, err

}
