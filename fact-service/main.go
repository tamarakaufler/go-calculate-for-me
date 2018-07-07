package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	factProto "github.com/tamarakaufler/go-calculate-for-me/pb/fact/v1"
	pingProto "github.com/tamarakaufler/go-calculate-for-me/pb/ping/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var port int

type server struct{}

var cache = map[uint64]uint64{}

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
	factProto.RegisterFactServiceServer(s, &server{})
	pingProto.RegisterPingServiceServer(s, &server{})
	reflection.Register(s)

	log.Printf("Starting Factorial Service server %s\n", host)
	if err := s.Serve(conn); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(ctx context.Context, r *factProto.FactRequest) (*factProto.FactResponse, error) {
	log.Printf("fact-service: Compute method: a=%d\n", r.A)
	a := factorial(r.A)

	return &factProto.FactResponse{Result: a}, nil
}

func (s *server) Ping(ctx context.Context, r *pingProto.PingRequest) (*pingProto.PingResponse, error) {
	pong := fmt.Sprintf("my pong to your %s", r.Ping)
	return &pingProto.PingResponse{Pong: pong}, nil
}

func factorial(a uint64) uint64 {
	log.Printf("fact-service: factorial function: a=%d\n", a)

	if a == 0 || a == 1 {
		return 1
	}

	b, ok := cache[a-1]
	key := a
	if ok {
		log.Printf("\tfact-service: factorial function: cache[%d]=%d\n", a-1, b)
		a = a * b
	} else {
		a = a * factorial(a-1)
	}
	log.Printf("\tfact-service: factorial function: caching fact for %d (%d)\n", key, a)
	cache[key] = a

	return a
}
