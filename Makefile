API_IMAGE_TAG=v1alpha1
GCD_IMAGE_TAG=v1alpha1
FACT_IMAGE_TAG=v1alpha1
FIB_IMAGE_TAG=v1alpha1
GCD_PORT?=3001
FACT_PORT?=3002
FIB_PORT?=3003
API_PORT?=3000

QUAY_PASS?=biggestsecret

net:
	docker network create --driver bridge calc-net-bridge

k8s:
	kubectl apply -f deployment/ -f deployment/prometheus/

protoc:
	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/tamarakaufler/go-calculate-for-me pb/gcd/v1/gcd.proto
	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/tamarakaufler/go-calculate-for-me pb/fact/v1/fact.proto
	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/tamarakaufler/go-calculate-for-me pb/fib/v1/fib.proto
	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/tamarakaufler/go-calculate-for-me pb/healtz/v1/healtz.proto
	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/tamarakaufler/go-calculate-for-me pb/ping/v1/ping.proto

dev-gcd-service: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gcd-service/gcd-service -a -installsuffix cgo gcd-service/main.go
	docker build -f gcd-service/Dockerfile -t quay.io/tamarakaufler/gcd-service:$(GCD_IMAGE_TAG) .

build-gcd-service: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gcd-service/gcd-service -a -installsuffix cgo gcd-service/main.go
	docker build -f gcd-service/Dockerfile -t quay.io/tamarakaufler/gcd-service:$(GCD_IMAGE_TAG) .
	docker login quay.io -u tamarakaufler -p $(QUAY_PASS)
	docker push quay.io/tamarakaufler/gcd-service:$(GCD_IMAGE_TAG)

run-gcd-service:
	docker run \
	--name=gcd-service \
	--network=calc-net-bridge \
	--rm \
	-d \
	-p $(GCD_PORT):3000 \
	quay.io/tamarakaufler/gcd-service:$(GCD_IMAGE_TAG) \
	-port=$(GCD_PORT)

dev-fact-service: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fact-service/fact-service -a -installsuffix cgo fact-service/main.go
	docker build -f fact-service/Dockerfile -t quay.io/tamarakaufler/factorial-service:$(FACT_IMAGE_TAG) .

build-fact-service: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fact-service/fact-service -a -installsuffix cgo fact-service/main.go
	docker build -f fact-service/Dockerfile -t quay.io/tamarakaufler/factorial-service:$(FACT_IMAGE_TAG) .
	docker login quay.io -u tamarakaufler -p $(QUAY_PASS)
	docker push quay.io/tamarakaufler/factorial-service:$(FACT_IMAGE_TAG)

run-fact-service:
	docker run \
	--name=fact-service \
	--network=calc-net-bridge \
	--rm \
	-d \
	-p $(FACT_PORT):3000 \
	quay.io/tamarakaufler/factorial-service:$(FACT_IMAGE_TAG) \
	-port=$(FACT_PORT)


dev-fib-service: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fib-service/fib-service -a -installsuffix cgo fib-service/main.go
	docker build -f fib-service/Dockerfile -t quay.io/tamarakaufler/fibonacci-service:$(FIB_IMAGE_TAG) .

build-fib-service: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fib-service/fib-service -a -installsuffix cgo fib-service/main.go
	docker build -f fib-service/Dockerfile -t quay.io/tamarakaufler/fibonacci-service:$(FIB_IMAGE_TAG) .
	docker login quay.io -u tamarakaufler -p $(QUAY_PASS)
	docker push quay.io/tamarakaufler/fibonacci-service:$(FIB_IMAGE_TAG)

run-fib-service:
	docker run \
	--name=fib-service \
	--network=calc-net-bridge \
	--rm \
	-d \
	-p $(FIB_PORT):3000 \
	quay.io/tamarakaufler/fibonacci-service:$(FIB_IMAGE_TAG) \
	-port=$(FIB_PORT)


dev-api-service: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-service/api-service -a -installsuffix cgo api-service/main.go
	docker build -f api-service/Dockerfile -t quay.io/tamarakaufler/api-calculations:$(API_IMAGE_TAG) .

build-api-service: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-service/api-service -a -installsuffix cgo api-service/main.go
	docker build -f api-service/Dockerfile -t quay.io/tamarakaufler/api-calculations:$(API_IMAGE_TAG) .
	docker login quay.io -u tamarakaufler -p $(QUAY_PASS)
	docker push quay.io/tamarakaufler/api-calculations:$(API_IMAGE_TAG)

run-api-service:
	docker run \
	--name=api-service \
	--network=calc-net-bridge \
	--rm \
	-p $(API_PORT):3000 \
	quay.io/tamarakaufler/api-calculations:$(API_IMAGE_TAG) \
	--gcd-port=$(GCD_PORT) --fact-port=$(FACT_PORT) --fib-port=$(FIB_PORT)


dev-all-services:  protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gcd-service/gcd-service -a -installsuffix cgo gcd-service/main.go
	docker build -f gcd-service/Dockerfile -t quay.io/tamarakaufler/gcd-service:$(GCD_IMAGE_TAG) .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fact-service/fact-service -a -installsuffix cgo fact-service/main.go
	docker build -f fact-service/Dockerfile -t quay.io/tamarakaufler/factorial-service:$(FACT_IMAGE_TAG) .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fib-service/fib-service -a -installsuffix cgo fib-service/main.go
	docker build -f fib-service/Dockerfile -t quay.io/tamarakaufler/fibonacci-service:$(FIB_IMAGE_TAG) .


dev-all: dev-all-services 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-service/api-service -a -installsuffix cgo api-service/main.go
	docker build -f api-service/Dockerfile -t quay.io/tamarakaufler/api-calculations:$(API_IMAGE_TAG) .

run-all-docker: run-gcd-service run-fact-service run-fib-service run-api-service

run-all-k8s: dev-all k8s

rmapicontainer:
	docker ps | grep "api-service" | awk '{print $1}' | xargs docker rm -f