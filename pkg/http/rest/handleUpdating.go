package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/williepotgieter/candyshop/pkg/updating"
)

func updateCandyName(us updating.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var ucc updating.CandyCategory
		var ucn updating.CandyName
		var ucp updating.CandyPrice
		params := mux.Vars(r)
		column := params["column"]

		id, err := uuid.Parse(params["id"])
		if err != nil {
			http.Error(w, "Invalid request params", http.StatusBadRequest)
			return
		}

		if column == "category" {
			if err := json.NewDecoder(r.Body).Decode(&ucc); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			if err := us.UpdateCandyCategory(id, ucc); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else if column == "name" {
			if err := json.NewDecoder(r.Body).Decode(&ucn); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			if err := us.UpdateCandyName(id, ucn); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else if column == "price" {
			if err := json.NewDecoder(r.Body).Decode(&ucp); err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			if err := us.UpdateCandyPrice(id, ucp); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{Code: http.StatusOK, Message: "Candy was updated"})
	}
}
