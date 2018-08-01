package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tamarakaufler/go-calculate-for-me/api-service/client"
	factProto "github.com/tamarakaufler/go-calculate-for-me/pb/fact/v1"
)

type FactOutput struct {
	Result uint64 `json:"result"`
}

func FactHandler(conf client.Config) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request for factorial calculation [%s]", r.RequestURI)

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

		// Call Fact service -----------------------------
		factClient, err := client.FactService(conf)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("error", http.StatusText(http.StatusInternalServerError))
			return
		}

		factReq := &factProto.FactRequest{A: a}
		if factRes, err := factClient.Compute(r.Context(), factReq); err == nil {
			w.WriteHeader(http.StatusOK)

			output := &FactOutput{
				Result: factRes.GetResult(),
			}
			json.NewEncoder(w).Encode(output)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("error", http.StatusText(http.StatusInternalServerError))
		}
	})
}
