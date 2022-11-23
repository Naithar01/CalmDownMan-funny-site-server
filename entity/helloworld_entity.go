package entity

import (
	"time"
)

type HelloWorld struct {
	Id         int       `json:"id"`
	World      string    `json:"world"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}
