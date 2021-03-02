package rest

import (
	"encoding/json"
	"net/http"

	"github.com/williepotgieter/candyshop/pkg/reading"
)

func getAllCandiesHandler(rs reading.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cs, err := rs.GetAllCandyNames()
		if err != nil {
			http.Error(w, "Cannot process your request at this time...", http.StatusInternalServerError)
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cs)
	}
}
