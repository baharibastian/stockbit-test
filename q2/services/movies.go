package services

import (
	"github.com/baharibastian/stockbit-test/q2/entity"
	"github.com/baharibastian/stockbit-test/q2/repositories"
)

type MoviesSvcInterface interface {
	Search(keyword, page string) (entity.SearchMoviesResult, error)
	GetByID(id string) (entity.SingleMovieResult, error)
}

type moviesService struct {
	moviesRepo repositories.MoviesRepo
}

func NewMoviesService(moviesRepo repositories.MoviesRepo) *moviesService {
	return &moviesService{
		moviesRepo: moviesRepo,
	}
}

func (ms moviesService) Search(keyword, page string) (entity.SearchMoviesResult, error) {
	if page == "" {
		page = "0"
	}
	return ms.moviesRepo.Search(keyword, page)
}

func (ms moviesService) GetByID(id string) (entity.SingleMovieResult, error) {
	return ms.moviesRepo.GetByID(id)
}
