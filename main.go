package main

import (
	"mygram-golang/conf"
	"mygram-golang/controller"
	"mygram-golang/middleware"
	"mygram-golang/repository"
	"mygram-golang/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := conf.InitDB()

	userRepository := repository.NewRepository(db)
	userService := service.NewService(userRepository)
	userController := controller.NewUserController(userService)

	router := gin.Default()

	// route
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.Login)
	router.PUT("/users", middleware.AuthMiddleware(), userController.UpdateUser)
	router.DELETE("/users", middleware.AuthMiddleware(), userController.DeleteUser)

	router.POST("/user/test", middleware.AuthMiddleware(), userController.TestUser)

	router.Run()
}
