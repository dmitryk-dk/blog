package models


type Post struct {
	Id			int	   `json:"id"`
	Title   	string `json:"title"`
	Description string `json:"description"`
}

func NewPost (id int, title, description string) *Post{
	return &Post{id, title, description}
}
