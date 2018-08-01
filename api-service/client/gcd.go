package client

import (
	"fmt"

	gcdProto "github.com/tamarakaufler/go-calculate-for-me/pb/gcd/v1"
	"google.golang.org/grpc"
)

type GCDClient struct {
	gcdProto.GCDServiceClient
}

func GCDService(config Config) (*GCDClient, error) {

	svc := fmt.Sprintf("%s:%d", config.Service, config.Port)
	conn, err := grpc.Dial(svc, config.Options...)
	if err != nil {
		return nil, err
	}
	gcdClient := gcdProto.NewGCDServiceClient(conn)

	return &GCDClient{gcdClient}, nil
}
