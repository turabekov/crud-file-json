package models

type Post struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      string `json:"user_id"`
}

type ResponseGetByIdPost struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	User        User   `json:"user"`
}
