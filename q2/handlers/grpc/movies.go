package grpc

import (
	"context"
	pb "github.com/baharibastian/stockbit-test/q2/pb/proto"
	"github.com/baharibastian/stockbit-test/q2/services"
)

type moviesGRPCHandler struct {
	pb.UnimplementedMoviesServiceServer
	service services.MoviesSvcInterface
}

func NewMoviesGRPC(moviesSvc services.MoviesSvcInterface) *moviesGRPCHandler {
	return &moviesGRPCHandler{
		service: moviesSvc,
	}
}

func (mc moviesGRPCHandler) Search(_ context.Context, request *pb.SearchMoviesRequest) (*pb.SearchMoviesResponse, error) {
	movies, err := mc.service.Search(request.GetSearchword(), string(request.GetPagination()))
	if err != nil {
		return nil, err
	}

	resp := &pb.SearchMoviesResponse{
		Search:       make([]*pb.MovieData, 0),
		Response:     movies.Response,
		TotalResults: movies.TotalResults,
	}

	for _, movie := range movies.Movies {
		resp.Search = append(resp.Search, &pb.MovieData{
			Title:  movie.Title,
			Year:   movie.Year,
			ImdbID: movie.IMDBId,
			Type:   movie.Type,
			Poster: movie.Poster,
		})
	}

	return resp, nil
}

func (mc moviesGRPCHandler) GetByID(_ context.Context, request *pb.GetByIDRequest) (*pb.GetByIDResponse, error) {
	movie, err := mc.service.GetByID(request.GetId())
	if err != nil {
		return nil, err
	}
	
	resp := &pb.GetByIDResponse{
		Title:      movie.Title,
		Year:       movie.Year,
		Rated:      movie.Rated,
		Released:   movie.Released,
		Runtime:    movie.Runtime,
		Genre:      movie.Genre,
		Director:   movie.Director,
		Writer:     movie.Writer,
		Actors:     movie.Actors,
		Plot:       movie.Plot,
		Language:   movie.Language,
		Country:    movie.Country,
		Awards:     movie.Awards,
		Poster:     movie.Poster,
		Ratings:    make([]*pb.Rating, 0),
		Metascore:  movie.Metascore,
		ImdbRating: movie.IMDBRating,
		ImdbVotes:  movie.IMDBVotes,
		ImdbID:     movie.IMDBId,
		Type:       movie.Type,
		DVD:        movie.DVD,
		BoxOffice:  movie.BoxOffice,
		Production: movie.Production,
		Website:    movie.Website,
		Response:   movie.Response,
	}

	for _, rating := range movie.Ratings {
		resp.Ratings = append(resp.Ratings, &pb.Rating{
			Source: rating.Source,
			Value:  rating.Value,
		})
	}

	return resp, nil
}