package main

import (
	"context"
	pb "github.com/baharibastian/stockbit-test/q2/pb/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	moviesConn := pb.NewMoviesServiceClient(conn)

	//mv, err := moviesConn.GetByID(context.Background(), &pb.GetByIDRequest{
	//	Id: "tt65512j4480",
	//})
	//if err != nil {
	//	log.Fatalf("Error when calling GetByID: %s", err)
	//}

	movies, err := moviesConn.Search(context.Background(), &pb.SearchMoviesRequest{
		Searchword: "spiderman",
		Pagination: 1,
	})

	log.Printf("Response from server: {%v}\n", movies)

}
