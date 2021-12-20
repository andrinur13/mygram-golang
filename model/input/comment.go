package input

type CommentInput struct {
	Message string `json:"message" binding:"required"`
	PhotoID int    `json:"photo_id"`
}

type DeleteComment struct {
	ID int `uri:"id" binding:"required"`
}

type UpdateComment struct {
	ID int `uri:"id" binding:"required"`
}
