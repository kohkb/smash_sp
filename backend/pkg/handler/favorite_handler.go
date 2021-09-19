package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kohkb/smash_sp/pkg/repository"
)

func GetFavorite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	favoriteRepository, err := repository.NewFavoriteRepositoryImpl()

	if err != nil {
		panic(err)
	}

	favorite, err := favoriteRepository.Find(vars["id"])

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(&favorite)
}

func GetFavorites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	favoriteRepository, err := repository.NewFavoriteRepositoryImpl()

	if err != nil {
		panic(err)
	}

	favorites, err := favoriteRepository.FindAll()

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(&favorites)
}

func CreateFavorite(w http.ResponseWriter, r *http.Request) {
	favoriteRepository, err := repository.NewFavoriteRepositoryImpl()

	if err != nil {
		panic(err)
	}

	favorite, err := favoriteRepository.Create(r.FormValue("video_id"))

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(&favorite)
}

func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	favoriteRepository, err := repository.NewFavoriteRepositoryImpl()

	if err != nil {
		panic(err)
	}

	err = favoriteRepository.Delete(vars["id"])

	if err != nil {
		panic(err)
	}
}
