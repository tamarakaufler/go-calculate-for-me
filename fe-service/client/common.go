package client

import (
	"google.golang.org/grpc"
)

type Config struct {
	Service string
	Port    int
	Options []grpc.DialOption
}
