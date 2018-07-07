FE_IMAGE_TAG=v1alpha1
GCD_IMAGE_TAG=v1alpha1
FACT_IMAGE_TAG=v1alpha1
FE_PORT?=3000
GCD_PORT?=3000
FACT_PORT?=3000

QUAY_PASS?=biggestsecret

net:
	docker network create --driver bridge calc-net-bridge

protoc:
	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/tamarakaufler/go-calculate-for-me pb/gcd/v1/gcd.proto
	protoc -I/usr/local/include -I. --go_out=plugins=grpc:$(GOPATH)/src/github.com/tamarakaufler/go-calculate-for-me pb/fact/v1/fact.proto
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
	-p $(GCD_PORT):$(GCD_PORT) \
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
	-p $(FACT_PORT):$(FACT_PORT) \
	quay.io/tamarakaufler/factorial-service:$(FACT_IMAGE_TAG) \
	-port=$(FACT_PORT)

dev-fe-service: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fe-service/fe-service -a -installsuffix cgo fe-service/main.go
	docker build -f fe-service/Dockerfile -t quay.io/tamarakaufler/fe-calculations:$(FE_IMAGE_TAG) .

build-fe-service: protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fe-service/fe-service -a -installsuffix cgo fe-service/main.go
	docker build -f fe-service/Dockerfile -t quay.io/tamarakaufler/fe-calculations:$(FE_IMAGE_TAG) .
	docker login quay.io -u tamarakaufler -p $(QUAY_PASS)
	docker push quay.io/tamarakaufler/fe-calculations:$(FE_IMAGE_TAG)

run-fe-service:
	docker run \
	--name=fe-service \
	--network=calc-net-bridge \
	--rm \
	-p $(FE_PORT):3000 \
	quay.io/tamarakaufler/fe-calculations:$(FE_IMAGE_TAG) \
	--gcd-port=$(GCD_PORT) --fact-port=$(FACT_PORT)


dev-all-services:  protoc
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gcd-service/gcd-service -a -installsuffix cgo gcd-service/main.go
	docker build -f gcd-service/Dockerfile -t quay.io/tamarakaufler/gcd-service:$(GCD_IMAGE_TAG) .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fact-service/fact-service -a -installsuffix cgo fact-service/main.go
	docker build -f fact-service/Dockerfile -t quay.io/tamarakaufler/factorial-service:$(FACT_IMAGE_TAG) .


dev-all: dev-all-services 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fe-service/fe-service -a -installsuffix cgo fe-service/main.go
	docker build -f fe-service/Dockerfile -t quay.io/tamarakaufler/fe-calculations:$(FE_IMAGE_TAG) .

run-all: run-gcd-service run-fact-service run-fe-service
