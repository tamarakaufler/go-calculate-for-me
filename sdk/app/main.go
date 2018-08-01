// An example how to use the client SDK for the calculate service
package main

import (
	"context"
	"flag"
	"log"
	"net/url"
	"os"
	"time"

	calcClient "github.com/tamarakaufler/go-calculate-for-me/sdk/calc-service/client"
)

var fibNum, factNum, gcdNum1, gcdNum2 uint64

func init() {
	flag.Uint64Var(&fibNum, "fib", 0, "Number for which Fibonacci shoud be calculated")
	flag.Uint64Var(&factNum, "fact", 0, "Number for which Factorial shoud be calculated")
	flag.Uint64Var(&gcdNum1, "gcd1", 0, "First number for which Global Common Denominator shoud be calculated")
	flag.Uint64Var(&gcdNum2, "gcd2", 0, "Second number for which Global Common Denominator shoud be calculated")
}

func main() {
	flag.Parse()

	u := &url.URL{
		Host:   "localhost:3000",
		Scheme: "http",
	}
	logger := log.New(os.Stdout, "calc-client: ", log.LstdFlags)
	cl := calcClient.NewClient(u, "calc-client", logger)

	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	if fibNum != 0 {
		out, err := cl.Fibonacci(ctx, fibNum)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Fibonacci result for %d is %d\n\n", fibNum, out.Result)
	}

	if factNum != 0 {
		out, err := cl.Factorial(ctx, factNum)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Factorial result for %d is %d\n\n", factNum, out.Result)
	}

	if gcdNum1 != 0 && gcdNum2 != 0 {
		out, err := cl.GreatestCommonDenominator(ctx, gcdNum1, gcdNum2)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("GreatestCommonDenominator result for %d and %d is %d\n\n", gcdNum1, gcdNum2, out.Result)
	}
	cancel()

}
