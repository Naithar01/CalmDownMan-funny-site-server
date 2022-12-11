package entity

import (
	"time"

	"github.com/Naithar01/CalmDownMan-funny-site-server/database"
)

type User struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}

type PostList_Author struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (p_a *PostList_Author) GetCategoryInfo(category_id int) error {
	err := database.Database.QueryRow("SELECT id, username FROM user WHERE id=?", category_id).Scan(&p_a.Id, &p_a.Username)

	if err != nil {
		return err
	}

	return nil
}
