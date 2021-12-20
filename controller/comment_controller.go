package controller

import (
	"mygram-golang/helper"
	"mygram-golang/model/input"
	"mygram-golang/model/response"
	"mygram-golang/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentService service.CommentService
	photoService   service.PhotoService
}

func NewCommentController(commentService service.CommentService, photoService service.PhotoService) *commentController {
	return &commentController{commentService, photoService}
}

func (h *commentController) AddNewComment(c *gin.Context) {
	var input input.CommentInput

	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessages := gin.H{
			"errors": errors,
		}

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// send to service
	newComment, err := h.commentService.CreateComment(input, currentUser)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)

		response := helper.APIResponse("failed", errorMessages)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCommentResponse := response.CreateCommentResponse{
		ID:      newComment.ID,
		Message: newComment.Message,
	}

	response := helper.APIResponse("created", newCommentResponse)
	c.JSON(http.StatusOK, response)
	return
}

func (h *commentController) DeleteComment(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "unauthorized user")
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	var idCommentUri input.DeleteComment

	err := c.ShouldBindUri(&idCommentUri)

	if err != nil {
		response := helper.APIResponse("failed", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	idComment := idCommentUri.ID

	if idComment == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	_, err = h.commentService.DeleteComment(idComment)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.APIResponse("ok", "success deleted comment!")
	c.JSON(http.StatusOK, response)
	return
}

func (h *commentController) GetComment(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var idPhotoUri input.UpdatePhotoIDUser

	err := c.ShouldBindUri(&idPhotoUri)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	comment, err := h.commentService.GetComment(currentUser)
	photo, err := h.photoService.GetPhotoByID(idPhotoUri.ID)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := helper.APIResponse("ok", response.GetAllComment(comment, photo))
	c.JSON(http.StatusOK, response)
	return
}

func (h *commentController) UpdateComment(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(int)

	if currentUser == 0 {
		response := helper.APIResponse("failed", "id must be exist!")
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	UpdateComment := input.CommentInput{}

	err := c.ShouldBindJSON(&UpdateComment)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	var idCommentUri input.UpdateComment

	err = c.ShouldBindUri(&idCommentUri)

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	id_comment := idCommentUri.ID

	queryResult, err := h.commentService.UpdateComment(id_comment, UpdateComment)

	if queryResult.ID == 0 {
		response := helper.APIResponse("failed", "comment not found!")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if err != nil {
		errorMessages := helper.FormatValidationError(err)
		response := helper.APIResponse("failed", gin.H{
			"errors": errorMessages,
		})
		c.JSON(http.StatusUnprocessableEntity, response)
	}

	Updated, err := h.photoService.GetPhotoByID(id_comment)

	response := helper.APIResponse("ok", Updated)
	c.JSON(http.StatusOK, response)
	return
}
