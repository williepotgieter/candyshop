package rest

import (
	"github.com/gorilla/mux"
	"github.com/williepotgieter/candyshop/pkg/adding"
	"github.com/williepotgieter/candyshop/pkg/reading"
)

func InitHandlers(rs reading.Service, as adding.Service) *mux.Router {
	router := mux.NewRouter()

	// Reading
	router.HandleFunc("/api/", welcomeHandler()).Methods("GET")
	router.HandleFunc("/api/candies", getAllCandiesHandler(rs)).Methods("GET")

	// Adding
	router.HandleFunc("/api/candy", addCandy(as)).Methods("POST")

	return router
}
