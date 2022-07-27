package model

type Blog struct {
	Id     int    `json:"Id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Url    string `json:"url"`
	Likes  int    `json:"likes"`
}
