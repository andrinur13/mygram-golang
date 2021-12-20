package response

import (
	"mygram-golang/model/entity"
	"time"
)

type CreateCommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"-"`
}

type GetCommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int       `json:"photo_id"`
	CreatedAt time.Time `json:"-"`
	Photo     entity.Photo
}

func GetAllComment(comments []entity.Comment, photo entity.Photo) []GetCommentResponse {
	if len(comments) == 0 {
		return []GetCommentResponse{}
	}

	var allComment []GetCommentResponse

	for _, comment := range comments {
		tmp := GetCommentResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			CreatedAt: comment.CreatedAt,
			Photo: entity.Photo{
				ID:       photo.ID,
				Title:    photo.Title,
				Caption:  photo.Caption,
				PhotoURL: photo.PhotoURL,
				UserID:   photo.UserID,
			},
		}

		allComment = append(allComment, tmp)
	}

	return allComment
}
