# go-calculate-for-me
A SAAS for calculating Greatest common denominator/Factorial/Fibonacci methods

The application runs as a suite of gRPC based microservices doing the calculations with a RESTful API service providing access to the functionality:

```
                         api-service
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
- gRPC
- gorilla
- Docker
- Kubernetes
- Prometheus
- Makefile

All code is stored and organised within a monorepo. Each service lives in its own directory. All protobuf descriptions share one directory (pb).The frontend to the calculation services (api-service) has (gorilla) handlers and (gRPC) clients stored in their respective subdirectories (handler, client). Kubernetes deployment yamls are stored in the deployment dir.

Makefile is used for ease of development and running.

The application is deployed in a Kubernetes cluster, in a calculations namespace.	

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
  GCD, factorial, fibonacci and API services must all listen on different ports. One possible setup:

    make dev-all
    
    GCD_PORT=4000 FACT_PORT=5000 FIB_PORT=6000 API_PORT=8888 make run-all

  where the API service is running on default port 3000 in the container but is exposed on port 8888 on the host. GCD, Factorial and Fibonacci services run and are exposed on port 4000, 5000 and 6000 respectively.

### In Kubernetes
kubectl apply -f deployment/

Then access the API service on:

  minikube service api-service -n calculations --url

eg,

    http://192.168.99.100:30831/ping
    
    http://192.168.99.100:30831/fib/10
    
    http://192.168.99.100:30831/fact/6
    
    http://192.168.99.100:30831/gcd/363/654


## Monitoring
The api service is instrumented for monitoring with Prometheus. The scraping
path is the default /metrics. Intrumentation middleware provides two custom
metrics:
  - requests_total
  - request_duration_milliseconds

### Load testing with Apache Bench
The application can be load tested using the Apache bench.

#### Apache Bench installation
On Ubuntu:

  sudo apt-get install apache2-utils

#### Create some sizeable traffic
Running against API service running in Kubernetes cluster. Your IP and port may differ.

ab -t 10 -n 10 http://192.168.99.100:30831/fib/5

ab -t 10 -n 20 http://192.168.99.100:30831/fib/10

ab -t 10 -n 30 http://192.168.99.100:30831/fib/100

ab -t 10 -n 50 http://192.168.99.100:30831/fact/5

ab -t 10 -n 60 http://192.168.99.100:30831/fact/10

ab -t 10 -n 70 http://192.168.99.100:30831/fact/100

ab -t 10 -n 100 http://192.168.99.100:30831/gcd/55/555

The above commands can be issued by running:

./load/ab_load.sh
