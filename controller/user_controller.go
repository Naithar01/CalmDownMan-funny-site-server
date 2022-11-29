package controller

import (
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity"
	"github.com/Naithar01/CalmDownMan-funny-site-server/entity/dto"
	"github.com/Naithar01/CalmDownMan-funny-site-server/service"
)

type UserController interface {
	GetAllUser() ([]entity.User, error)
	CreateUser(user dto.CreateUserDto) (int, error)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userSerivce service.UserService) UserController {
	return &userController{
		userService: userSerivce,
	}
}

func (u userController) GetAllUser() ([]entity.User, error) {
	return u.userService.GetAllUser()
}

// Register
func (u userController) CreateUser(user dto.CreateUserDto) (int, error) {
	return u.userService.CreateUser(user)
}
