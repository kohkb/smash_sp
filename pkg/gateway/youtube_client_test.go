package gateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewYoutubeClient(t *testing.T) {
	client := NewYoutubeClient()
	assert.Equal(t, "https://www.googleapis.com/youtube/v3", client.BaseURL)
}

// func TestSearchVideosByQuery(t *testing.T) {
// 	client := NewYoutubeClient()

// 	_, err := client.SearchVideosByQuery("hoge")

// 	assert.Nil(t, err)
// }
