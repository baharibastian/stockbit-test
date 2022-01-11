package q2

import (
	grpcHandler "github.com/baharibastian/stockbit-test/q2/handlers/grpc"
	pb "github.com/baharibastian/stockbit-test/q2/pb/proto"
	"github.com/baharibastian/stockbit-test/q2/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RunGRPCServer(port string, moviesSvc services.MoviesSvcInterface) {
	ls, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterMoviesServiceServer(grpcServer, grpcHandler.NewMoviesGRPC(moviesSvc))

	log.Println("GRPC Server Listening on Port :"+port)
	if err := grpcServer.Serve(ls); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %s: %v", port, err)
	}
}