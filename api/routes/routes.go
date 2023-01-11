package routes

import (
	"github.com/kipropfredrick/DockBlog/api/controller"
	"github.com/kipropfredrick/DockBlog/infrastructure"

)

// PostRoute -> Route for question module
type PostRoute struct {
	Controller controller.PostController
	UserController controller.UerServiceController
	Handler    infrastructure.GinRouter
}

// NewPostRoute -> initializes new choice rouets
func NewPostRoute(
	controller controller.PostController,
	UserController controller.UerServiceController,
	handler infrastructure.GinRouter,

) PostRoute {
	return PostRoute{
		Controller: controller,
		UserController:UserController,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p PostRoute) Setup() {
	allroutes := p.Handler.Gin.Group("/posts") //Router group
	{
		allroutes.POST("/register",p.UserController.CreateUser)
		allroutes.POST("/login",p.UserController.Login)
		allroutes.GET("/", p.Controller.GetPosts)
		allroutes.POST("/", p.Controller.AddPost)
		allroutes.GET("/:id", p.Controller.GetPost)
		allroutes.DELETE("/:id", p.Controller.DeletePost)
		allroutes.PUT("/:id", p.Controller.UpdatePost)
	    
		//
	}
}
