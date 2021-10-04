package main

import (
	"log"
	"net"

	pb "github.com/incidrthreat/gohost/service"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGoHostServer
	Alive bool
}

func (s *server) IsAlive(ctx context.Context, in *pb.AliveRequest) {
	log.Printf("Data Received from: %v", in.GetHostname())
	s.Alive = in.GetAlive()
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGoHostServer{s, &server{}}
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
