package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/incidrthreat/gohost/gohost"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGoHostClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	hostname, _ := os.Hostname()

	r, err := c.IsAlive(ctx, &pb.AliveRequest{Hostname: hostname, Alive: true})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Accepted: %v", r.GetResponse())
}
