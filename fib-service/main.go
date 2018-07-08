package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	fibProto "github.com/tamarakaufler/go-calculate-for-me/pb/fib/v1"
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
	fibProto.RegisterFibServiceServer(s, &server{})
	pingProto.RegisterPingServiceServer(s, &server{})
	reflection.Register(s)

	log.Printf("Starting Fibonacci Service server %s\n", host)
	if err := s.Serve(conn); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(ctx context.Context, r *fibProto.FibRequest) (*fibProto.FibResponse, error) {
	log.Printf("fib-service: Compute method: a=%d\n", r.A)
	a := fibonacci(r.A)

	return &fibProto.FibResponse{Result: a}, nil
}

func (s *server) Ping(ctx context.Context, r *pingProto.PingRequest) (*pingProto.PingResponse, error) {
	pong := fmt.Sprintf("my pong to your %s", r.Ping)
	return &pingProto.PingResponse{Pong: pong}, nil
}

func fibonacci(a uint64) uint64 {
	log.Printf("fib-service: fibonacci function: a=%d\n", a)

	if a == 0 {
		cache[0] = 0
		return 0
	}
	if a == 1 {
		cache[1] = 1
		return 1
	}

	key := a
	var value uint64
	b, ok1 := cache[a-1]
	c, ok2 := cache[a-2]

	if !ok1 {
		log.Printf("\tfib-service: fibonacci function: no cache[%d]\n", a-1)
		if !ok2 {
			log.Printf("\tfib-service: fibonacci function: no cache[%d]\n", a-2)
			value = fibonacci(a-1) + fibonacci(a-2)
		} else {
			log.Printf("\tfib-service: fibonacci function: cache[%d]=%d\n", a-2, c)
			value = fibonacci(a-1) + c
		}
	} else {
		if !ok2 {
			log.Printf("\tfib-service: fibonacci function: no cache[%d]\n", a-2)
			value = b + fibonacci(a-2)
		} else {
			log.Printf("\tfib-service: fibonacci function: cache[%d]=%d\n", a-2, c)
			value = b + c
		}
	}
	log.Printf("\tfib-service: fibonacci function: caching fib for %d (%d)\n", key, value)
	cache[key] = value

	return value
}
