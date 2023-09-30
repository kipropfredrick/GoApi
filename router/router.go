package router

import (
	"server/internal/users"
	"server/util"
)

type PostRoute struct {
	UserHandler users.Handler
	handler     GinRouter
}

//func to hanlde all routes

func NewUserRoute(
	uHandler users.Handler,
	handler GinRouter,
) PostRoute {
	return PostRoute{
		UserHandler: uHandler,
		handler:     handler,
	}
}

//set up route function
func (p PostRoute) Setup() {
	public := p.handler.Gin.Group("api/user")
	{
		public.GET("/home", p.UserHandler.Home)
		public.POST("/login/", p.UserHandler.LoginUser)
		public.POST("/create", p.UserHandler.NewUserCreate)
		public.POST("/register", p.UserHandler.CreateNewUsers)
		public.POST("/validate", p.UserHandler.ValidateAcc)
	}
	protected := p.handler.Gin.Group("/api")
	{
		protected.Use(util.JwtAuthMiddleware())
		protected.GET("/logout/", p.UserHandler.Logout)
	}
}
