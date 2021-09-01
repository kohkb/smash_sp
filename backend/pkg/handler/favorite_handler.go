package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kohkb/smash_sp/pkg/repository"
)

func GetFavorites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// TODO: Remove after vue setting
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
