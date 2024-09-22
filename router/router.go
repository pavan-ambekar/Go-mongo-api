package router

import (
	"github.com/gorilla/mux"
	"github.com/pavan-ambekar/Go-mongo-api/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controllers.DeleteAllMovies).Methods("DELETE")

	return router
}
