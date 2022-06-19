package routes

import (
	"api/controllers"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return RestigerRoutes(r)
}

func RestigerRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/", controllers.GetPost).Methods("GET")
	return r
}
