package repositories

import (
	"encoding/json"
	"errors"
	"github.com/baharibastian/stockbit-test/q2/entity"
	"log"
	"net/http"
)

type moviesRepo struct {
	baseURL string
	apiKey  string
}

type MoviesRepo interface {
	Search(keyword, page string) (entity.SearchMoviesResult, error)
	GetByID(id string) (entity.SingleMovieResult, error)
}

func NewMovies(baseURL, apiKey string) *moviesRepo {
	return &moviesRepo{
		baseURL: baseURL,
		apiKey: apiKey,
	}
}

func (mr moviesRepo) Search(keyword, page string) (entity.SearchMoviesResult, error) {
	client := &http.Client{}

	var result entity.SearchMoviesResult

	request, err := http.NewRequest(http.MethodGet, mr.baseURL, nil)
	if err != nil {
		return result, err
	}
	q := request.URL.Query()
	q.Add("apikey", mr.apiKey)
	q.Add("s", keyword)
	q.Add("page", page)
	request.URL.RawQuery = q.Encode()

	log.Println("Request URL", request.URL)

	response, err := client.Do(request)
	if err != nil {
		return result, err
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return result, err
	}

	if result.Response == "False" {
		return result, errors.New(result.Error)
	}

	return result, nil
}

func (mr moviesRepo) GetByID(id string) (entity.SingleMovieResult, error) {
	client := &http.Client{}

	var result entity.SingleMovieResult

	request, err := http.NewRequest(http.MethodGet, mr.baseURL, nil)
	if err != nil {
		return result, err
	}
	q := request.URL.Query()
	q.Add("apikey", mr.apiKey)
	q.Add("i", id)
	request.URL.RawQuery = q.Encode()

	response, err := client.Do(request)
	if err != nil {
		return result, err
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return result, err
	}

	if result.Response == "False" {
		return result, errors.New(result.Error)
	}

	return result, nil
}
