package application

import (

	//"github.com/kevinkimutai/metadata/internal/adapter/db/db"

	"github.com/kevinkimutai/metadata/internal/app/core/domain"
	"github.com/kevinkimutai/metadata/internal/ports"
)

type MovieRepo struct {
	db ports.MovieRepoPort
}

func NewMovieRepo(db ports.MovieRepoPort) *MovieRepo {
	return &MovieRepo{db: db}
}

// Movies
func (m *MovieRepo) CreateNewMovie(movie domain.Movie) (domain.Movie, error) {

	movie, err := m.db.CreateMovie(movie)

	return movie, err
}

func (m *MovieRepo) GetMovies(movieParams domain.MovieParams) (domain.FetchData, error) {

	data, err := m.db.GetAllMovies(movieParams)

	return data, err
}

func (m *MovieRepo) GetMovie(movieID int64) (*domain.Movie, error) {
	//TODO:HANDLE ERRORS
	movie, err := m.db.GetMovieById(movieID)

	return movie, err
}
