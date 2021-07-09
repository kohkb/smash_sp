package gateway

type YoutubeClient struct {
	Name string
}

func (yc YoutubeClient) Method1() string {
	return yc.Name
}
