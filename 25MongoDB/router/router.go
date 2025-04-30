package router

import (
	"github.com/gorilla/mux"
	"github.com/ketan-chaudhary/mongodb/controller"
)

func Router() *mux.Router {
	routr := mux.NewRouter()

	routr.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	routr.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	routr.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	routr.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	routr.HandleFunc("/api/deleteall", controller.DeleteAllMovie).Methods("DELETE")

	return routr
}
