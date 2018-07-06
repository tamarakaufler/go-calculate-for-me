# go-calculate-for-me
Running communicating gRPC based microservices in Kubernetes

## Protocol buffers
Autogenerate grpc code by running the following commands in the root derectory:

	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$GOPATH/src/github.com/tamarakaufler/go-calculate-for-me pb/gcd/v1/gcd.proto
  
	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$GOPATH/src/github.com/tamarakaufler/go-calculate-for-me pb/fact/v1/fact.proto

	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$GOPATH/src/github.com/tamarakaufler/go-calculate-for-me pb/healtz/v1/healtz.proto

	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$GOPATH/src/github.com/tamarakaufler/go-calculate-for-me pb/ping/v1/ping.proto