package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kohkb/smash_sp/pkg/gateway"
)

func GetVideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := gateway.NewYoutubeClient()
	videos, err := client.SearchVideosByQuery(r.URL.Query().Get("q"))

	if err != nil {
		panic("runtime error")
	}

	json.NewEncoder(w).Encode(&videos)
}
