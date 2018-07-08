package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tamarakaufler/go-calculate-for-me/fe-service/client"
	"github.com/tamarakaufler/go-calculate-for-me/fe-service/handler"
	"google.golang.org/grpc"
)

var (
	port, gcdPort, factPort, fibPort int
)

func init() {
	flag.IntVar(&port, "port", 3000, "FE port")
	flag.IntVar(&gcdPort, "gcd-port", 3000, "GCD service port")
	flag.IntVar(&factPort, "fact-port", 3000, "Factorial service port")
	flag.IntVar(&fibPort, "fib-port", 3000, "Fibonacci service port")
}

func main() {
	flag.Parse()
	r := mux.NewRouter()

	fmt.Printf("Starting FE service on port %d\n", port)

	gcdConf := client.Config{
		Service: "gcd-service",
		Port:    gcdPort,
		Options: []grpc.DialOption{grpc.WithInsecure()},
	}
	gcdHandler := handler.GCDHandler(gcdConf)

	factConf := client.Config{
		Service: "fact-service",
		Port:    factPort,
		Options: []grpc.DialOption{grpc.WithInsecure()},
	}
	factHandler := handler.FactHandler(factConf)

	fibConf := client.Config{
		Service: "fib-service",
		Port:    fibPort,
		Options: []grpc.DialOption{grpc.WithInsecure()},
	}
	fibHandler := handler.FibHandler(fibConf)

	r.HandleFunc("/gcd/{a}/{b}", gcdHandler).
		Methods("GET").
		Name("gcd-compute")
	r.HandleFunc("/fact/{a}", factHandler).
		Methods("GET").
		Name("fact-compute")
	r.HandleFunc("/fib/{a}", fibHandler).
		Methods("GET").
		Name("fib-compute")
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"alive": true}`)
	}).
		Methods("GET").
		Name("ping")

	host := fmt.Sprintf(":%d", port)
	log.Printf("Starting FE server %s\n", host)
	log.Fatal(http.ListenAndServe(host, r))
}
