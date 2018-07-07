package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	gcdProto "github.com/tamarakaufler/go-calculate-for-me/pb/gcd/v1"
	pingProto "github.com/tamarakaufler/go-calculate-for-me/pb/ping/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

var port int

func init() {
	flag.IntVar(&port, "port", 3000, "Port on which the RPC server is listening")
}

func main() {
	flag.Parse()

	host := fmt.Sprintf(":%d", port)
	conn, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	gcdProto.RegisterGCDServiceServer(s, &server{})
	pingProto.RegisterPingServiceServer(s, &server{})
	reflection.Register(s)

	log.Printf("Starting Greatest common denominator Service server %s\n", host)
	if err := s.Serve(conn); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(ctx context.Context, r *gcdProto.GCDRequest) (*gcdProto.GCDResponse, error) {
	log.Printf("gcd-service: Compute method: a=%d, b=%d\n", r.A, r.B)
	a, b := r.A, r.B
	for b != 0 {
		a, b = b, a%b
	}
	return &gcdProto.GCDResponse{Result: a}, nil
}

func (s *server) Ping(ctx context.Context, r *pingProto.PingRequest) (*pingProto.PingResponse, error) {
	pong := fmt.Sprintf("my pong to your %s", r.Ping)
	return &pingProto.PingResponse{Pong: pong}, nil
}
