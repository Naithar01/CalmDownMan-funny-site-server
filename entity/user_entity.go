package entity

import "time"

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
