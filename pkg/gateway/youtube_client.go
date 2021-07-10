package gateway

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type YoutubeClient struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

// TODO 別のパッケージへ移動させる
type Video struct {
	Id string
}

type VideoList struct {
	Videos []Video
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

func (c *YoutubeClient) SearchVideosByQuery(query string) (*VideoList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/search?key=%s&q=%s&order=viewCount", c.BaseURL, c.apiKey, query), nil)

	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	bodyString := string(body)

	fmt.Println(bodyString)

	return nil, err
}
