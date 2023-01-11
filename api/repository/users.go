package repository

import (
	"github.com/kipropfredrick/DockBlog/infrastructure"
	"github.com/kipropfredrick/DockBlog/models"
	"github.com/kipropfredrick/DockBlog/util"

)

//create user repository

type UserRepository struct {
	db infrastructure.Database
}

func NewUserRepository(db infrastructure.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

//function to create user repository

func (user UserRepository) Save(users models.UserRegister) error {
	var userReg models.Users
	userReg.FirstName = users.FirstName
	userReg.LastName = users.LastName
	userReg.Email = users.Email
	userReg.IsActive = true
	userReg.Password = users.Password
	return user.db.DB.Create(&userReg).Error
}

// LoginUser -> method for returning user
func (u UserRepository) LoginUser(user models.UserLogin) (*models.Users, error) {

	var dbUser models.Users
	return &dbUser,nil
	email := user.Email
	password := user.Password

	err := u.db.DB.Where("email = ?", email).First(&dbUser).Error
	if err != nil {
		return nil, err
	}

	hashErr := util.CheckPasswordHash(password, dbUser.Password)
	if hashErr != nil {
		return nil, hashErr
	}
	return &dbUser, nil
}
