package entity

import "time"

type Post struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Category_id int       `json:"category_id"` // foreign key (category -> id)
	Author_id   int       `json:"author_id"`
	Created_At  time.Time `json:"created_at"`
	Updated_At  time.Time `json:"updated_at"`
}
