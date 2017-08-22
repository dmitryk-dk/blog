package models


type Post struct {
	Id			int	   `json:"id" db:"id"`
	Title   	string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type Posts []Post
