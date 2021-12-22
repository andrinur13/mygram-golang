package main

import (
	"log"
	"mygram-golang/conf"
	"mygram-golang/controller"
	"mygram-golang/middleware"
	"mygram-golang/repository"
	"mygram-golang/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := conf.InitDB()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	photoController := controller.NewPhotoController(photoService, userService)

	socialmediaRepository := repository.NewSocialMediaRepository(db)
	socialmediaService := service.NewSocialMediaService(socialmediaRepository)
	socialmediaController := controller.NewSocialMediaController(socialmediaService, userService)

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository)
	commentController := controller.NewCommentController(commentService, photoService)

	router := gin.Default()

	// route
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.Login)
	router.PUT("/users", middleware.AuthMiddleware(), userController.UpdateUser)
	router.DELETE("/users", middleware.AuthMiddleware(), userController.DeleteUser)

	// photos
	router.POST("/photos", middleware.AuthMiddleware(), photoController.AddNewPhoto)
	router.DELETE("/photos/:id", middleware.AuthMiddleware(), photoController.DeletePhoto)
	router.GET("/photos", middleware.AuthMiddleware(), photoController.GetPhotos)
	router.GET("/photos/:id", photoController.GetPhoto)
	router.PUT("/photos/:id", middleware.AuthMiddleware(), photoController.UpdatePhoto)

	// comments
	router.POST("/comments", middleware.AuthMiddleware(), commentController.AddNewComment)
	router.DELETE("/comments/:id", middleware.AuthMiddleware(), commentController.DeleteComment)
	router.GET("/comments", middleware.AuthMiddleware(), commentController.GetComment)
	router.PUT("/comments/:id", middleware.AuthMiddleware(), commentController.UpdateComment)

	// social media
	router.POST("/socialmedias", middleware.AuthMiddleware(), socialmediaController.AddNewSocialMedia)
	router.GET("/socialmedias", middleware.AuthMiddleware(), socialmediaController.GetSocialMedia)
	router.PUT("/socialmedias/:id", middleware.AuthMiddleware(), socialmediaController.UpdateSocialMedia)
	router.DELETE("/socialmedias/:id", middleware.AuthMiddleware(), socialmediaController.DeleteSocialmedia)

	router.POST("/user/test", middleware.AuthMiddleware(), userController.TestUser)

	// data := input.UpdatePhoto{
	// 	Title:    "title 1",
	// 	Caption:  "Title 1 adalah",
	// 	PhotoURL: "https;//facebook.com",
	// }

	// photoService.UpdatePhoto(8, data)

	router.Run()
}
