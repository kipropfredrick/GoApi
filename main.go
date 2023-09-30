package main

import (
	"fmt"
	"log"
	"server/db"
	"server/internal/repository"
	"server/internal/servicesa"
	"server/internal/testin"
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
	testRepo := testin.NewTestrepo(dbConn.GetDB())
	userService := servicesa.NewUserService(userRepo)
	testService := testin.NewTestservice(testRepo)
	testHnadl := testin.NewTestHanlder(testService)
	userHandler := users.NewHandler(userService, testService)
	userPost := router.NewUserRoute(*userHandler, *testHnadl, routers)
	userPost.Setup()
	// initialize the router and
	// router.InitRouter(userHandler)
	// router.Start("0.0.0.0:8080")
	routers.Gin.Run(":3005")
	fmt.Println("Connecting to database")
}
