package gateway

import "fmt"

type YoutubeClient struct {
	Name string
}

func (yc YoutubeClient) Method1() {
	fmt.Println(yc.Name)
}
