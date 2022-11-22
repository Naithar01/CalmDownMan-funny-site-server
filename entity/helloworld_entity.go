package entity

import (
	"time"
)

type HelloWorld struct {
	Id        int       `json:"id"`
	World     string    `json:"world"`
	Create_At time.Time `json:"create_at"`
	Delete_At time.Time `json:"delete_at"`
}
