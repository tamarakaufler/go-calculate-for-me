package handler

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tamarakaufler/go-calculate-for-me/fe-service/client"
	gcdProto "github.com/tamarakaufler/go-calculate-for-me/pb/gcd/v1"
)

func GCDHandler(conf client.Config) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		bStr := vars["b"]
		b, err := strconv.ParseUint(bStr, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("error", "bad [b] value")
			return
		}

		// Call GCD service -----------------------------
		gcdClient, err := client.GCDService(conf)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		pbReq := &gcdProto.GCDRequest{A: a, B: b}
		if pbRes, err := gcdClient.Compute(r.Context(), pbReq); err == nil {
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, fmt.Sprintf(`{"result": "%d"}`, pbRes.GetResult()))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("error", "bad [b] value")
		}
	})
}
