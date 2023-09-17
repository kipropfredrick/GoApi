package main

import (
	"github.com/kipropfredrick/DockBlog/infrastructure"
  "github.com/kipropfredrick/DockBlog/api/routes"
	"github.com/kipropfredrick/DockBlog/api/service"
  "github.com/kipropfredrick/DockBlog/api/repository"
  "github.com/kipropfredrick/DockBlog/models"
	"github.com/kipropfredrick/DockBlog/api/controller"

)

func init() {
	infrastructure.LoadEnv()
}
func main() {
	router := infrastructure.NewGinRouter() //router has been initialized and configured
    db := infrastructure.NewDatabase() // databse has been initialized and configured
    postRepository := repository.NewPostRepository(db) // repository are being setup
    userRepository := repository.NewUserRepository(db)
    postService := service.NewPostService(postRepository) // service are being setup
    UerServiceController := service.CreateUser(userRepository)
    postController := controller.NewPostController(postService) // controller are being set up
    userControll := controller.UserControllers(UerServiceController)
    postRoute := routes.NewPostRoute(postController,userControll, router) // post routes are initialized
    postRoute.Setup() // post routes are being setup
    db.DB.AutoMigrate(&models.Post{},&models.Users{}) // migrating Post model to datbase table
    router.Gin.Run(":8005") //server started on 8000 port
}
