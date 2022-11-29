package service

import (
	"github.com/Naithar01/CalmDownMan-funny-site-server/action"
	"github.com/Naithar01/CalmDownMan-funny-site-server/database"
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity"
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity/dto"
)

type UserService interface {
	GetAllUser() ([]entity.User, error)
	CreateUser(user dto.CreateUserDto) (int, error)
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

func (u userService) CreateUser(user dto.CreateUserDto) (int, error) {
	hash_pass, err := action.UserRegisterHashPassWord(user.Password)

	if err != nil {
		return 0, err
	}

	c_user, err := database.Database.Exec("INSERT INTO user (username, password) VALUES (?, ?)", user.Username, hash_pass)

	if err != nil {
		return 0, err
	}

	Id, err := c_user.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(Id), nil
}
