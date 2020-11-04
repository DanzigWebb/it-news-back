package models

type Post struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Img         string `json:"img"`
	Description string `json:"description"`
}

type HabrPost struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	ImgUrl      string   `json:"img"`
	Published   string   `json:"published"`
	Categories  []string `json:"categories"`
}
