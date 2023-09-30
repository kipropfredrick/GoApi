package models

import (
	"context"
)

type User struct {
	ID       int64  `json:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName string `json:"last_name" db:"last_name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Phone string `json:"phone" db:"phone"`
	ID_NO string `json:"id_number" db:"id_numner"`
	PlateNumber string `json:"plate_number" db:"plate_numner"`
	LicenseNumber string `json:"license_number" db:"license_numner"`
	Image string `json:"image_caption" db:"image_caption"`
	IsVerified bool  `json:"is_verified" db:"is_verified"`
	Test  []TestIn 

}
// test struct
type TestIn struct {
	ID   int64  `json:"id"`
	NAME string `json:"name"`
	Users string `json:"user_id" db:"user_id"`
}
//user request
type CreateUserRequest struct{
	FirstName string `json:"first_name" db:"first_name"`
	LastName string `json:"last_name" db:"last_name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Phone string `json:"phone" db:"phone"`
	ID_NO string `json:"id_number" db:"id_numner"`
	PlateNumber string `json:"plate_number" db:"plate_numner"`
	LicenseNumber string `json:"license_number" db:"license_numner"`
	Image string `json:"image_caption" db:"image_caption"`
	IsVerified bool  `json:"is_verified" db:"is_verified"`
	Hashed_code string  `json:"hashed_code" db:"hashed_code"`
	Verify string  `json:"verify" db:"verify"`
	ID int64 `json:"id" db:"id"`

}
//create user response
type CreateUserResponse struct {
	ID       string  `json:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName string `json:"last_name" db:"last_name"`
    Phone string `json:"phone" db:"phone"`
	ID_NO string `json:"id_number" db:"id_numner"`
	Email    string `json:"email" db:"email"`
	Test  []TestIn 
	Code string `json:"code" db:"code"`
	IsVerified bool  `json:"is_verified" db:"is_verified"`

}
//login user response struct
type LoginUserResponse struct {
	Accesstoken string
	ID       string  `json:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName string `json:"last_name" db:"last_name"`
    Phone string `json:"phone" db:"phone"`
	Email string `json:"email" db:"email"`
	ID_NO string `json:"id_number" db:"id_numner"`
	Test  []TestIn 

}
//login user response struct
type LoginUserRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
//user repository struct
type Repository interface {
	CreateUser(ctx context.Context,hashed_code string, user *User) (*User, error)
	GetByEmail(ctx context.Context,email string)(*User,error)
	Check_exist_email(ctx context.Context,user_id string)(*User,error)
	UpdateUser(ctx context.Context,user_id string,user *User) (*User, error)
	GetCode(ctx context.Context,uid int64)(*User,error)
     Verify(ctx context.Context,user_id int64) (*CreateUserResponse, error)
}
// user service interface
type Service interface {
	CreateNewUser(ctx context.Context,req *CreateUserRequest)(*CreateUserResponse, error)
	UpdateUser(ctx context.Context,req *CreateUserRequest)(*CreateUserResponse, error)
	ValidateAcc(ctx context.Context,req *CreateUserRequest)(*CreateUserResponse, error)
	Login(c context.Context,req *LoginUserRequest)(*LoginUserResponse,error)
}