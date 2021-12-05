package controller

import (
	"mygram-golang/helper"
	"mygram-golang/model/input"
	"mygram-golang/model/response"
	"mygram-golang/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var input input.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessages := gin.H{"errors": err}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send to service
	newUser, err := h.userService.CreateUser(input)

	if err != nil {
		errorMessages := gin.H{"errors": err}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUserResponse := response.UserRegisterResponse{
		ID:        newUser.ID,
		Age:       newUser.Age,
		Email:     newUser.Email,
		Password:  newUser.Password,
		Username:  newUser.Username,
		CreatedAt: newUser.CreatedAt,
	}

	response := helper.APIResponse("created", newUserResponse)
	c.JSON(http.StatusOK, response)
	return
}
