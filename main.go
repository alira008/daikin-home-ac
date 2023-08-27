package main

import (
	"log"
    "daikin-home-ac/server"
)

func main() {
    address := ":5520"
	httpServer := server.NewHttpServer(address)

	log.Printf("Starting server on %s\n", address)
	log.Fatal(httpServer.ListenAndServe())
}

