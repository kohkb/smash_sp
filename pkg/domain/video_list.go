package domain

type VideoList struct {
	Videos []Video
}

func (vl *VideoList) ToUrl() []string {
	var urls []string
	for _, video := range vl.Videos {
		urls = append(urls, video.Url())
	}
	return urls
}
