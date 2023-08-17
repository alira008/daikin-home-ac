package server

import (
    "net/http"
    "github.com/gorilla/mux"
)

type httpServer struct {

}

func NewHttpServer(addr string) *http.Server {
    server := &httpServer{

    }

    r:= mux.NewRouter()
    r.HandleFunc("/temperature/{temp}", server.postTemperature).Methods("GET")
    r.HandleFunc("/mode/{state}", postTemperature).Methods("GET")
    r.HandleFunc("/timer/{state}", postTemperature).Methods("GET")
    r.HandleFunc("/timerdelay/{minutes}", postTemperature).Methods("GET")
    r.HandleFunc("/power/{state}", postTemperature).Methods("GET")
    r.HandleFunc("/fanspeed/{state}", postTemperature).Methods("GET")
    r.HandleFunc("/swing/{state}", postTemperature).Methods("GET")
    r.HandleFunc("/powerful", postTemperature).Methods("GET")
    r.HandleFunc("/econo", postTemperature).Methods("GET")

    return &http.Server{
        Addr: ":5520",
        Handler: r,
    }
}

