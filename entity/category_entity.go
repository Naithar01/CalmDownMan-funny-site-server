package entity

import (
	"time"

	"github.com/Naithar01/CalmDownMan-funny-site-server/database"
)

type Category struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

type PostList_Category struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func (p_c *PostList_Category) GetCategoryInfo(category_id int) error {
	err := database.Database.QueryRow("SELECT id, title FROM category WHERE id=?", category_id).Scan(&p_c.Id, &p_c.Title)

	if err != nil {
		return err
	}

	return nil
}
