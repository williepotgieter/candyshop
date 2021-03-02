package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/williepotgieter/candyshop/pkg/adding"
)

func addCandy(as adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var nc adding.Candy

		if err := json.NewDecoder(r.Body).Decode(&nc); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		id, err := as.AddCandy(nc)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(struct {
			ID uuid.UUID `json:"id"`
		}{ID: id})
	}
}
