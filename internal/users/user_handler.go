package users

import (
	// "fmt"
	"fmt"
	"net/http"
	"server/internal/models"
	"server/internal/testin"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	models.Service
	testin.TestService
}

// create handler function to create
func NewHandler(s models.Service,t testin.TestService) *Handler {
	return &Handler{
		Service: s,
	}
}

// function to cfeate user
func (h *Handler) NewUserCreate(g *gin.Context) {
    var u models.CreateUserRequest
	
	if err := g.ShouldBindJSON(&u); err != nil {
		g.JSON(http.StatusInternalServerError,gin.H{"err":err.Error()})
		return
	}
	res,_:=h.Service.UpdateUser(g.Request.Context(),&u)

	fmt.Println("done",res);
	g.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"message":"succefuly UPDATE","data":res})
}
// 
func (h *Handler) CreateNewUsers(g *gin.Context) {
    var u models.CreateUserRequest
	
	if err := g.ShouldBindJSON(&u); err != nil {
		g.JSON(http.StatusInternalServerError,gin.H{"err":err.Error()})
		return
	}
	res,_:=h.Service.CreateNewUser(g.Request.Context(),&u)

	fmt.Println("done",res);
	g.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"message":"succefuly created","data":res})
}
// validate account
func (h *Handler) ValidateAcc(g *gin.Context) {
    var u models.CreateUserRequest
	
	if err := g.ShouldBindJSON(&u); err != nil {
		g.JSON(http.StatusInternalServerError,gin.H{"err":err.Error()})
		return
	}
	res,err:=h.Service.ValidateAcc(g.Request.Context(),&u)
	if err != nil {
		g.JSON(http.StatusExpectationFailed,gin.H{"status":http.StatusExpectationFailed,"message":"Failed to Validate Acount plz contact admin for assistance","data":res})
	   return
	}
	// else if(res.IsVerified) {
	// 	g.JSON(int(http.StatusAccepted),gin.H{"status":http.StatusAccepted,"message":"Your Account is already verified","data":res})
	// 	return	
	// }

	fmt.Println("done",res);
	g.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"message":"Account succesfully verified","data":res})
}
//login handler
func (h *Handler) LoginUser(g *gin.Context) {
	var  user models.LoginUserRequest
if err := g.ShouldBindJSON(&user);err != nil {
	g.JSON(http.StatusBadRequest,gin.H{"internal server error":err.Error()})
}

u,err := h.Service.Login(g.Request.Context(),&user)
if err != nil {
	g.JSON(http.StatusInternalServerError,gin.H{"error loging in":err.Error()})
}

// g.SetCookie("jwt",u.Accesstoken,3600,"/","localhost",false,true)
resp :=&models.LoginUserResponse{
	Accesstoken: u.Accesstoken,
	ID: u.ID,
	FirstName: u.FirstName,
	LastName: u.LastName,
	Phone: u.Phone,
	ID_NO: u.ID_NO,
	Test: u.Test,
	Email: u.Email,
}
// g.JSON(http.StatusOK,resp)
g.JSON(http.StatusOK,gin.H{"status":http.StatusOK,"message":"succefuly login","data":resp})
}
//function to logout
func (h *Handler) Logout(g *gin.Context) {
   g.SetCookie("jwt","",-1,"","",false,false)
   g.JSON(http.StatusOK,gin.H{"message": "logout successful"})
}
func (h *Handler) Home(g *gin.Context) {
	// g.SetCookie("jwt","",-1,"","",false,false)
	g.JSON(http.StatusOK,gin.H{"message": "endpoint working"})
 }