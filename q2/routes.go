package q2

import (
	httpHandler "github.com/baharibastian/stockbit-test/q2/handlers/http"
	"github.com/baharibastian/stockbit-test/q2/services"
	"log"
	"net/http"
)

func RunServer(port string, moviesSvc services.MoviesSvcInterface) {
	moviesHandler := httpHandler.NewMoviesHandler(moviesSvc)
	http.HandleFunc("/movies/search", moviesHandler.SearchMovies)
	http.HandleFunc("/movie", moviesHandler.GetByID)

	log.Println("HTTP Server listening on port :"+port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
