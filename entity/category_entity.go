package entity

import "time"

type Category struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}
