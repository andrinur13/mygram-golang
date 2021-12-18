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

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	photoController := controller.NewPhotoController(photoService, userService)

	router := gin.Default()

	// route
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.Login)
	router.PUT("/users", middleware.AuthMiddleware(), userController.UpdateUser)
	router.DELETE("/users", middleware.AuthMiddleware(), userController.DeleteUser)

	router.POST("/photos", middleware.AuthMiddleware(), photoController.AddNewPhoto)
	router.DELETE("/photos/:id", middleware.AuthMiddleware(), photoController.DeletePhoto)
	router.GET("/photos", middleware.AuthMiddleware(), photoController.GetPhotos)
	router.GET("/photos/:id", photoController.GetPhoto)
	router.PUT("/photos/:id", middleware.AuthMiddleware(), photoController.UpdatePhoto)

	router.POST("/user/test", middleware.AuthMiddleware(), userController.TestUser)

	// data := input.UpdatePhoto{
	// 	Title:    "title 1",
	// 	Caption:  "Title 1 adalah",
	// 	PhotoURL: "https;//facebook.com",
	// }

	// photoService.UpdatePhoto(8, data)

	router.Run()
}
