package service

import (
	"github.com/kipropfredrick/DockBlog/api/repository"
	"github.com/kipropfredrick/DockBlog/models"

)

//create struct of type service

type UserService struct {
	userRepo repository.UserRepository
}

//create a function that take userservice as an argument and return

func CreateUser(userRepo repository.UserRepository) UserService {
	return UserService{
		userRepo: userRepo,
	}
}

//registratio service

func (userService UserService) RegisterUser(userReg models.UserRegister) error {
	return userService.userRepo.Save(userReg)
}

//function to login

func (userService UserService) LoginUsers(userReg models.UserLogin) (*models.Users, error){
	return userService.userRepo.LoginUser(userReg)
}