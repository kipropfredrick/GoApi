package servicesa

import (
	"context"
	"fmt"
	"log"
	"os"
	"server/internal/models"
	"server/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

//to be moved to env file
const(
	secret="sasasassa"
)
//type struct thataccept repository as arguement
type UserService struct {
	models.Repository
	timeout time.Duration
}

//function constructor that usersivercie struct as arguement and
func NewUserService(repository models.Repository) models.Service {
	// return pointer to user service struct
    return &UserService{
		repository,
		time.Duration(2)*time.Second,
	}
}
//jwt claims struct
type MyjwtClaims struct {
	ID string `json:"id"`
	LastName string `json:"last_name"`
	Name string `json:"name"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	jwt.RegisteredClaims
}
//function to create user 
func (s *UserService) CreateNewUser(c context.Context,req *models.CreateUserRequest)(*models.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	fmt.Println("hased",hashedPassword)
	if err != nil {
		return nil, err
	}
	// create hashed code
	var randomstr=util.RanndStrin(8)
	hashed_code, err := util.HashPassword(strconv.Itoa(randomstr))
	fmt.Println("hased",hashed_code)
	if err != nil {
		return nil, err
	}
    req.Hashed_code=hashed_code
	req.Password=hashedPassword
	u := &models.User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Phone: req.Phone,
		ID_NO: req.ID_NO,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx,hashed_code,u)
	if err != nil {
		return nil, err
	}
    
	res:=&models.CreateUserResponse{
		ID:strconv.Itoa(int(r.ID)),//type cast to string
		Email: r.Email,
		FirstName: req.FirstName,
		LastName: req.LastName,
		Phone: req.Phone,
		ID_NO: req.ID_NO,
		Test:r.Test,
		Code: strconv.Itoa(randomstr),
	}

	return res, nil

}
// update user
func (s *UserService) UpdateUser(c context.Context,req *models.CreateUserRequest)(*models.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	fmt.Println("hased",hashedPassword)
	if err != nil {
		return nil, err
	}

	u := &models.User{
		LicenseNumber: req.LicenseNumber,
		Phone: req.Phone,
		ID_NO: req.ID_NO,
		PlateNumber:    req.PlateNumber,
		Password: hashedPassword,
	}

	r, err := s.Repository.UpdateUser(ctx,req.Phone,u)
	if err != nil {
		return nil, err
	}
    
	res:=&models.CreateUserResponse{
		ID:strconv.Itoa(int(r.ID)),//type cast to string
		Email: r.Email,
		FirstName: req.PlateNumber,
		LastName: req.LicenseNumber,
		Phone: req.Phone,
		ID_NO: req.ID_NO,
		Test:r.Test,
	}

	return res, nil

}
//login user
func (s UserService) Login(c context.Context,req *models.LoginUserRequest)(*models.LoginUserResponse,error){
ctx ,cancel:=context.WithTimeout(c,s.timeout)
defer cancel()


u,err :=s.Repository.GetByEmail(ctx,req.Email)
if err != nil {
	return &models.LoginUserResponse{},err
}

hashedPassword, _ := util.HashPassword(req.Password)
fmt.Println("u",u,hashedPassword)
err= util.CheckPassword(req.Password,u.Password)
if err != nil {
	return &models.LoginUserResponse{}, err
}
//generate jwt if password matches the database
  token :=jwt.NewWithClaims(jwt.SigningMethodHS256,MyjwtClaims{
	ID: strconv.Itoa(int(u.ID)),
	FirstName: u.FirstName,
	LastName: u.LastName,
	Email: u.Email,
	RegisteredClaims: jwt.RegisteredClaims{
		Issuer: strconv.Itoa(int(u.ID)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*time.Hour)),
	
	},
  })
  signedstring,err :=token.SignedString([]byte(os.Getenv("API_SECRET")))
  if err != nil {
	return &models.LoginUserResponse{}, err
  }
		fmt.Println("service",u.Test);
		lores:=&models.LoginUserResponse{ 
			Accesstoken:signedstring,
			ID:strconv.Itoa(int(u.ID)),
			FirstName: u.FirstName,
			LastName: u.LastName,
			Phone:u.Phone,
			ID_NO: u.ID_NO,
			Test:u.Test, 
			Email: u.Email,
			}
			return lores,nil
}
// validate user and verfify
func (s UserService) ValidateAcc(c context.Context,req *models.CreateUserRequest)(*models.CreateUserResponse,error){
	ctx ,cancel:=context.WithTimeout(c,s.timeout)
	defer cancel()
	
	
	code,err :=s.Repository.GetCode(ctx,req.ID)
	if err != nil || code.IsVerified == true {
		// log.Fatal(err)
		// fmt.Println("Error",err)
		return &models.CreateUserResponse{IsVerified: code.IsVerified},err
	}
	
	hashedPassword, _ := util.HashPassword(req.Verify)
	fmt.Println("u",code,hashedPassword)
	err= util.CheckPassword(req.Verify,code.Password)
	if err != nil {
		log.Fatal(err)
		fmt.Println("error",err)
		return nil, err
	}
	//generate jwt if password matches the database
	 
	verify, err := s.Repository.Verify(ctx,req.ID,)
	if err != nil {
		return nil, err
	}

	return verify, nil
	}