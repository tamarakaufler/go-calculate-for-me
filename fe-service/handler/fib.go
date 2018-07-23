package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tamarakaufler/go-calculate-for-me/fe-service/client"
	fibProto "github.com/tamarakaufler/go-calculate-for-me/pb/fib/v1"
)

func FibHandler(conf client.Config) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request for fibonacci calculation [%s]", r.RequestURI)

		w.Header().Set("Content-Type", "application/json")

		// Process input -------------------------------
		vars := mux.Vars(r)
		aStr := vars["a"]
		a, err := strconv.ParseUint(aStr, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("content", "")
			w.Header().Set("error", "bad [a] value")
			return
		}

		// Call Fib service -----------------------------
		fibClient, err := client.FibService(conf)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fibReq := &fibProto.FibRequest{A: a}
		if fibRes, err := fibClient.Compute(r.Context(), fibReq); err == nil {
			w.WriteHeader(http.StatusOK)

			io.WriteString(w, fmt.Sprintf(`{"result": "%d"}`, fibRes.GetResult()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("error", http.StatusText(http.StatusInternalServerError))
		}
	})
}
