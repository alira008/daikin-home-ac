package main

import (
	"fmt"
	"log"
	"net/http"
    "daikin-home-ac/internal/server"

	"github.com/gorilla/mux"
)

func postTemperature(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Temperature: %s", vars["temp"])
}

func main() {
	httpServer := server.NewHttpServer(":5040")
	log.Println("Starting server")
	log.Fatal(httpServer.ListenAndServe())
}

