package main

import (
	"log"
	"net/http"

	"github.com/kohkb/smash_sp/pkg/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/videos/search", handler.GetVideos).Methods("GET")
	r.HandleFunc("/videos/favorites/{id:[0-9]+}", handler.GetFavorite).Methods("GET")
	r.HandleFunc("/videos/favorites", handler.GetFavorites).Methods("GET")
	r.HandleFunc("/videos/favorites", handler.CreateFavorite).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
