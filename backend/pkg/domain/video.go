package domain

import (
	"time"
)

type Video struct {
	Id          string
	Title       string
	PublishTime time.Time
}

func (v *Video) Url() string {
	return "https://www.youtube.com/watch?v=" + v.Id
}

type Videos struct {
	Videos []Video
}

func (vl *Videos) ToUrl() []string {
	var urls []string
	for _, video := range vl.Videos {
		urls = append(urls, video.Url())
	}
	return urls
}
