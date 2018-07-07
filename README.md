# go-calculate-for-me
Running interdependent gRPC based microservices in Kubernetes

Application consists of several microservices, bulk of which implement and provide a result of a particular calculation, and one acting as an access point/api to the calculation services:
```
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
```

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

  ### Locally using Docker containers

  GCD, factorial and FE services must all listen on different ports. One possible setup:

    GCD_PORT=4000 FACT_PORT=5000 FE_PORT=5000 FE_PORT=8888 make run-fe-service

  where the FE service is running on port 3000 in the container but is exposed on port 8888 on the host. GCD and Factorial services run and are exposed on port 4000 and 5000 respectively.

  ### In Kubernetes

kubectl apply -f deployment/

Then access the FE service on:

  minikube service fe-service --url

eg,

    http://192.168.99.100:31298/ping
    
    http://192.168.99.100:31298/fact/6
    
    http://192.168.99.100:31298/gcd/363/654
