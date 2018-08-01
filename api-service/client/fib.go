package client

import (
	"fmt"

	fibProto "github.com/tamarakaufler/go-calculate-for-me/pb/fib/v1"
	"google.golang.org/grpc"
)

type FibClient struct {
	fibProto.FibServiceClient
}

func FibService(config Config) (*FibClient, error) {

	svc := fmt.Sprintf("%s:%d", config.Service, config.Port)
	conn, err := grpc.Dial(svc, config.Options...)
	if err != nil {
		return nil, err
	}
	factClient := fibProto.NewFibServiceClient(conn)

	return &FibClient{factClient}, nil
}
