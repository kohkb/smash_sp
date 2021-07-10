package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/kohkb/smash_sp/pkg/domain"
)

type YoutubeClient struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

const (
	BaseURL = "https://www.googleapis.com/youtube/v3"
)

func NewYoutubeClient() *YoutubeClient {
	return &YoutubeClient{
		BaseURL: BaseURL,
		apiKey:  os.Getenv("YOUTUBE_API_KEY"),
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *YoutubeClient) SearchVideosByQuery(query string) (*domain.VideoList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/search?key=%s&q=%s&order=viewCount", c.BaseURL, c.apiKey, query), nil)

	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var data map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	var video_list domain.VideoList

	for _, item := range data["items"].([]interface{}) {
		var video domain.Video
		video.Id = item.(map[string]interface{})["id"].(map[string]interface{})["videoId"].(string)

		video_list.Videos = append(video_list.Videos, video)
	}

	return &video_list, err
}
