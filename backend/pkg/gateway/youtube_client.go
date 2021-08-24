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

func (c *YoutubeClient) SearchVideosByQuery(query string) (*domain.Videos, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/search?key=%s&q=%s&order=viewCount&part=snippet", c.BaseURL, c.apiKey, query), nil)

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

	var videos domain.Videos

	for _, row := range data["items"].([]interface{}) {
		var video domain.Video

		item := row.(map[string]interface{})

		video.Id = item["id"].(map[string]interface{})["videoId"].(string)
		video.Title = item["snippet"].(map[string]interface{})["title"].(string)
		publishTime, _ := time.Parse("2006-01-02T15:04:05Z", item["snippet"].(map[string]interface{})["publishTime"].(string))
		video.PublishTime = publishTime

		videos.Videos = append(videos.Videos, video)
	}

	return &videos, err
}
