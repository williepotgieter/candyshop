package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/williepotgieter/candyshop/pkg/adding"
	"github.com/williepotgieter/candyshop/pkg/deleting"
	"github.com/williepotgieter/candyshop/pkg/http/rest"
	"github.com/williepotgieter/candyshop/pkg/reading"
	"github.com/williepotgieter/candyshop/pkg/storage"
	"github.com/williepotgieter/candyshop/pkg/updating"
)

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	r, err := storage.SetupStorage()
	if err != nil {
		log.Fatalln("error while setting up storage", err)
	}

	// Adding
	as := adding.NewService(r)

	// Reading
	rs := reading.NewService(r)

	// Updating
	us := updating.NewService(r)

	// Deleting
	ds := deleting.NewService(r)

	fmt.Println("starting server on port ", os.Getenv("PORT"))
	router := rest.InitHandlers(as, rs, us, ds)
	log.Fatal(http.ListenAndServe(port, router))
}
