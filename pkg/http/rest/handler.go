package rest

import (
	"github.com/gorilla/mux"
	"github.com/williepotgieter/candyshop/pkg/adding"
	"github.com/williepotgieter/candyshop/pkg/reading"
	"github.com/williepotgieter/candyshop/pkg/updating"
)

func InitHandlers(as adding.Service, rs reading.Service, us updating.Service) *mux.Router {
	router := mux.NewRouter()

	// Adding
	router.HandleFunc("/api/candy", addCandy(as)).Methods("POST")

	// Reading
	router.HandleFunc("/api/candies", getAllCandiesHandler(rs)).Methods("GET")

	// Updating
	router.HandleFunc("/api/candy/{id}/{column}", updateCandyName(us)).Methods("PATCH")

	return router
}
