package gateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHoge(t *testing.T) {
	var youtubeClient YoutubeClient
	youtubeClient.Name = "hoge"
	assert.Equal(t, "hoge", youtubeClient.Method1())
}
