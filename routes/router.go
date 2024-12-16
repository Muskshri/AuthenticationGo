package routes

import (
	"gomod/controllers"

	"github.com/gorilla/mux"
)

func SetRouter(router *mux.Router) {
	router.HandleFunc("/regiser", controller.Register).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
}