package main

import (
	"fmt"
	"log"
	"server/db"
	"server/internal/repository"
	"server/internal/servicesa"
	"server/internal/users"
	"server/router"
)

func main() {
	routers := router.NewGinRouter()
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Couldnt connect to db %s", err)
	}
	// Create
	userRepo := repository.NewUserRepository(dbConn.GetDB())
	userService := servicesa.NewUserService(userRepo)

	userHandler := users.NewHandler(userService)
	userPost := router.NewUserRoute(*userHandler, routers)
	userPost.Setup()
	// initialize the router and
	// router.InitRouter(userHandler)
	// router.Start("0.0.0.0:8080")
	routers.Gin.Run(":3005")
	fmt.Println("Connecting to database")
}
