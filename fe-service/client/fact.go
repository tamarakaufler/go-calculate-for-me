package client

import (
	"fmt"

	factProto "github.com/tamarakaufler/go-calculate-for-me/pb/fact/v1"
	"google.golang.org/grpc"
)

type FactClient struct {
	factProto.FactServiceClient
}

func FactService(config Config) (*FactClient, error) {

	svc := fmt.Sprintf("%s:%d", config.Service, config.Port)
	conn, err := grpc.Dial(svc, config.Options...)
	if err != nil {
		return nil, err
	}
	factClient := factProto.NewFactServiceClient(conn)

	return &FactClient{factClient}, nil
}
