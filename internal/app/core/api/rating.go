package application

import (

	//"github.com/kevinkimutai/metadata/internal/adapter/db/db"

	"github.com/kevinkimutai/metadata/internal/app/core/domain"
	"github.com/kevinkimutai/metadata/internal/ports"
)

type RatingRepo struct {
	db ports.RatingRepoPort
}

func NewRatingRepo(db ports.RatingRepoPort) *RatingRepo {
	return &RatingRepo{db: db}
}

// Ratings
func (r *RatingRepo) CreateNewRating(rating domain.Rating) (domain.Rating, error) {

	rating, err := r.db.CreateRating(rating)

	return rating, err
}
