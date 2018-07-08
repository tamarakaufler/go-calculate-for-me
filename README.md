# go-calculate-for-me
Running interdependent gRPC based microservices in Kubernetes

Application consists of an API service (fe-service) and a collection of microservices, implementing a particular calculation:
```
                         fe-service
          (REST api to gRPC calculation microservices)
                              |
                              |
                -------------------------------
                |               |             |
                |               |             |
                |               |             |
                |               |             |
          gcd-service           |       fib-service
  (Greatest common denominator) |       (Fibonacci)
                                |
                      fact-service
                       (Factorial)
```
#### Technology stack and tools
- Golang
- microservices
- protocol buffers
- gRpc
- gorilla
- Docker
- Kubernetes
- Makefile

All code is stored and organised within a monorepo. Each service lives in its own directory. All protobuf descriptions share one directory (pb).The frontend to the calculation services (fe-service) has (gorilla) handlers and (gRPC) clients stored in their respective subdirectories (handler, client). Kubernetes deployment yamls are housed together in the deployment dir.

Makefile is used for ease of development and running.

## Protocol buffers
Autogenerate grpc code by running the following commands in the root derectory:

a) Manually

	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$GOPATH/src/github.com/tamarakaufler/go-calculate-for-me pb/gcd/v1/gcd.proto
  
	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$GOPATH/src/github.com/tamarakaufler/go-calculate-for-me pb/fact/v1/fact.proto
  
	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$GOPATH/src/github.com/tamarakaufler/go-calculate-for-me pb/fib/v1/fib.proto

	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$GOPATH/src/github.com/tamarakaufler/go-calculate-for-me pb/healtz/v1/healtz.proto

	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$GOPATH/src/github.com/tamarakaufler/go-calculate-for-me pb/ping/v1/ping.proto

b) make protoc

  ## Deployment

  ### Locally using Docker containers

  GCD, factorial, fibonacci and FE services must all listen on different ports. One possible setup:

    make dev-all
    
    GCD_PORT=4000 FACT_PORT=5000 FIB_PORT=6000 FE_PORT=8888 make run-all

  where the FE service is running on default port 3000 in the container but is exposed on port 8888 on the host. GCD, Factorial and Fibonacci services run and are exposed on port 4000, 5000 and 6000 respectively.

  ### In Kubernetes

kubectl apply -f deployment/

Then access the FE service on:

  minikube service fe-service --url

eg,

    http://192.168.99.100:31298/ping
    
    http://192.168.99.100:31298/fib/10
    
    http://192.168.99.100:31298/fact/6
    
    http://192.168.99.100:31298/gcd/363/654
