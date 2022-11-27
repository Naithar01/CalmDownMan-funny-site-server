package service

import (
	"github.com/Naithar01/CalmDownMan-funny-site-server/database"
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity"
)

type UserService interface {
	GetAllUser() ([]entity.User, error)
}

type userService struct {
	users []entity.User
}

func NewUserService() UserService {
	return &userService{
		users: []entity.User{},
	}
}

func (u userService) GetAllUser() ([]entity.User, error) {
	rows, err := database.Database.Query("SELECT * FROM user")

	var users []entity.User

	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var user entity.User

		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Created_At, &user.Updated_At)

		users = append(users, user)
	}

	return users, nil
}
