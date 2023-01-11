package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/kipropfredrick/DockBlog/api/service"
	"github.com/kipropfredrick/DockBlog/models"
	"github.com/kipropfredrick/DockBlog/util"

)

///creta controller service

type UerServiceController struct {
	userController service.UserService
}

//create an instance to userController serveice as an arguement

func UserControllers(useControll service.UserService) UerServiceController {
	return UerServiceController{
		userController: useControll,
	}
}

//function to sent data to userregister service

// CreateUser ->  calls CreateUser services for validated user
func (u *UerServiceController) CreateUser(c *gin.Context) {
	var user models.UserRegister
	if err := c.ShouldBind(&user); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Inavlid Json Provided")
		return
	}

	hashPassword, _ := util.HashPassword(user.Password)
	user.Password = hashPassword

	err := u.userController.RegisterUser(user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to create user")
		return
	}

	util.SuccessJSON(c, http.StatusOK, "Successfully Created user")
}

// LoginUser : Generates JWT Token for validated user
func (u *UerServiceController) Login(c *gin.Context) {
	var user models.UserLogin
	
	var hmacSampleSecret []byte
	if err := c.ShouldBindJSON(&user); err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Inavlid Json Provideddd")
		return
	}
	dbUser, err := u.userController.LoginUsers(user)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Invalid Login Credentials")
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": dbUser,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		util.ErrorJSON(c, http.StatusBadRequest, "Failed to get token")
		return
	}
	response := &util.Response{
		Success: true,
		Message: "Token generated sucessfully",
		Data:    tokenString,
	}
	c.JSON(http.StatusOK, response)
}
