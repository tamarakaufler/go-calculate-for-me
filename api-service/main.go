package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tamarakaufler/go-calculate-for-me/api-service/client"
	"github.com/tamarakaufler/go-calculate-for-me/api-service/handler"
	"github.com/tamarakaufler/go-calculate-for-me/instrumentation/monitoring"
	"google.golang.org/grpc"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	port, gcdPort, factPort, fibPort int
)

func init() {
	flag.IntVar(&port, "port", 3000, "API port")
	flag.IntVar(&gcdPort, "gcd-port", 3000, "GCD service port")
	flag.IntVar(&factPort, "fact-port", 3000, "Factorial service port")
	flag.IntVar(&fibPort, "fib-port", 3000, "Fibonacci service port")
}

func main() {
	flag.Parse()
	r := mux.NewRouter()

	promMiddler := monitoring.NewMonitMiddler("fe")

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

	r.Handle("/gcd/{a}/{b}", promMiddler.Instrument(gcdHandler)).
		Methods("GET").
		Name("gcd-compute")
	r.Handle("/fact/{a}", promMiddler.Instrument(factHandler)).
		Methods("GET").
		Name("fact-compute")
	r.Handle("/fib/{a}", promMiddler.Instrument(fibHandler)).
		Methods("GET").
		Name("fib-compute")

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"alive": true}`)
	}).
		Methods("GET").
		Name("ping")
	r.Handle("/metrics", promhttp.Handler()).
		Methods("GET").
		Name("metrics")

	host := fmt.Sprintf(":%d", port)
	log.Printf("Starting API server %s\n", host)
	log.Fatal(http.ListenAndServe(host, r))
}
