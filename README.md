# go-calculate-for-me
Running interdependent gRPC based microservices in Kubernetes

Application consists of several microservices, bulk of which implement and provide a result of a particular calculation, and one acting as an access point/api to the calculation services:

                 fe-service
  (REST api to gRPC calculation microservices)
                      |
                      |
                -------------
                |           |
                |           |
          gcd-service       |
  (Greatest common denominator)
                            |
                      fact-service
                       (Factorial)

- Golang
- microservices
- protocol buffers
- gRpc
- gorilla
- Docker
- Kubernetes
- Makefile

## Protocol buffers
Autogenerate grpc code by running the following commands in the root derectory:

	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$GOPATH/src/github.com/tamarakaufler/go-calculate-for-me pb/gcd/v1/gcd.proto
  
	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$GOPATH/src/github.com/tamarakaufler/go-calculate-for-me pb/fact/v1/fact.proto

	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$GOPATH/src/github.com/tamarakaufler/go-calculate-for-me pb/healtz/v1/healtz.proto

	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$GOPATH/src/github.com/tamarakaufler/go-calculate-for-me pb/ping/v1/ping.proto

  ## Deployment
