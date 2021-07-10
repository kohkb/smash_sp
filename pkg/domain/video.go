package domain

type Video struct {
	Id string
}

func (v *Video) Url() string {
	return "https://www.youtube.com/watch?v=" + v.Id
}
