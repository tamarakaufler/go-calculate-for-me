package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tamarakaufler/go-calculate-for-me/api-service/client"
	gcdProto "github.com/tamarakaufler/go-calculate-for-me/pb/gcd/v1"
)

type GCDOutput struct {
	Result uint64 `json:"result"`
}

func GCDHandler(conf client.Config) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request for greatest common denominator calculation [%s]", r.RequestURI)

		w.Header().Set("Content-Type", "application/json")

		// Process input -------------------------------
		vars := mux.Vars(r)
		aStr := vars["a"]
		a, err := strconv.ParseUint(aStr, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("content", "")
			w.Header().Set("error", "bad request ([a] value)")
			return
		}
		bStr := vars["b"]
		b, err := strconv.ParseUint(bStr, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("error", "bad request ([b] value)")
			return
		}

		// Call GCD service -----------------------------
		gcdClient, err := client.GCDService(conf)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("error", http.StatusText(http.StatusInternalServerError))
			return
		}

		gcdReq := &gcdProto.GCDRequest{A: a, B: b}
		if gcdRes, err := gcdClient.Compute(r.Context(), gcdReq); err == nil {
			w.WriteHeader(http.StatusOK)

			output := &GCDOutput{
				Result: gcdRes.GetResult(),
			}
			json.NewEncoder(w).Encode(output)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("error", http.StatusText(http.StatusInternalServerError))
		}
	})
}
