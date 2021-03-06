package main

import (
	"log"
	"net/http"

	"github.com/kohkb/smash_sp/pkg/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/videos/search", handler.GetVideos).Methods("GET")
	r.HandleFunc("/api/videos/favorites/{id:[0-9]+}", handler.GetFavorite).Methods("GET")
	r.HandleFunc("/api/videos/favorites/{id:[0-9]+}", handler.DeleteFavorite).Methods("DELETE")
	r.HandleFunc("/api/videos/favorites", handler.GetFavorites).Methods("GET")
	r.HandleFunc("/api/videos/favorites", handler.CreateFavorite).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))
}
