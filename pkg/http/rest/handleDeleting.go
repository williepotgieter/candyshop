package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/williepotgieter/candyshop/pkg/deleting"
)

func deleteCandy(ds deleting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, err := uuid.Parse(params["id"])
		if err != nil {
			http.Error(w, "Invalid request params", http.StatusBadRequest)
			return
		}

		if err := ds.DeleteCandy(id); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{Code: http.StatusOK, Message: "Candy was deleted"})
	}
}
