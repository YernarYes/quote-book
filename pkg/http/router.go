package httpserver

import (
	"quotes/internal/core/handler"

	"github.com/gorilla/mux"
)

func RoutesRegister(handler *handler.Handler, router *mux.Router) {
	router.HandleFunc("/quotes/random", handler.GetRandom).Methods("GET")
	router.HandleFunc("/quotes", handler.Create).Methods("POST")
	router.HandleFunc("/quotes", handler.Get).Methods("GET")
	router.HandleFunc("/quotes/{id}", handler.Delete).Methods("DELETE")
}
