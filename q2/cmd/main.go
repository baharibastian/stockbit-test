package main

import (
	"github.com/baharibastian/stockbit-test/q2"
	"github.com/baharibastian/stockbit-test/q2/config"
	"github.com/baharibastian/stockbit-test/q2/repositories"
	"github.com/baharibastian/stockbit-test/q2/services"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	moviesRepo := repositories.NewMovies(cfg.BASEUrl, cfg.APIKey)
	moviesSvc := services.NewMoviesService(moviesRepo)

	go q2.RunGRPCServer(cfg.GRPCPort, moviesSvc)
	q2.RunServer(cfg.HTTPPort, moviesSvc)
}
