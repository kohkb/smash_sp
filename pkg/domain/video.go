package domain

import "time"

type Video struct {
	Id          string
	Title       string
	PublishTime time.Time
}

func (v *Video) Url() string {
	return "https://www.youtube.com/watch?v=" + v.Id
}
