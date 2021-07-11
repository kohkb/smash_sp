package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrl(t *testing.T) {
	var video Video
	video.Id = "videoid"
	assert.Equal(t, "https://www.youtube.com/watch?v=videoid", video.Url())

}
