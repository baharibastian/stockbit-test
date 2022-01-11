package http

import (
	"encoding/json"
	"github.com/baharibastian/stockbit-test/q2/services"
	"log"
	"net/http"
)

type moviesHttpHandler struct {
	service services.MoviesSvcInterface
}

func NewMoviesHandler(service services.MoviesSvcInterface) *moviesHttpHandler {
	return &moviesHttpHandler{
		service: service,
	}
}

func (mh moviesHttpHandler) SearchMovies(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("pagination")
	keyword := r.URL.Query().Get("searchword")

	resp := make(map[string]interface{})

	movies, err := mh.service.Search(keyword, page)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		resp["message"] = err.Error()
		jsonResp, _ := json.Marshal(resp)
		log.Printf("[SearchMovies] err: %v", err.Error())
		w.Write(jsonResp)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp["message"] = "Success"
	resp["data"] = movies
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
	return
}

func (mh moviesHttpHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	imdbID := r.URL.Query().Get("id")

	resp := make(map[string]interface{})

	movie, err := mh.service.GetByID(imdbID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		resp["message"] = err.Error()
		jsonResp, _ := json.Marshal(resp)
		log.Printf("[SearchMovies] err: %v", err.Error())
		w.Write(jsonResp)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp["message"] = "Success"
	resp["data"] = movie
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
	return
}