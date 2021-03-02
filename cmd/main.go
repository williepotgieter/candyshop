package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/williepotgieter/candyshop/pkg/adding"
	"github.com/williepotgieter/candyshop/pkg/http/rest"
	"github.com/williepotgieter/candyshop/pkg/reading"
	"github.com/williepotgieter/candyshop/pkg/storage"
)

func main() {

	r, err := storage.SetupStorage()
	if err != nil {
		log.Fatalln("error while setting up storage", err)
	}

	// Reading
	rs := reading.NewService(r)

	// Adding
	as := adding.NewService(r)

	fmt.Println("starting server on port 5000")
	router := rest.InitHandlers(rs, as)
	log.Fatal(http.ListenAndServe(":5000", router))
}
