package response

import (
	"mygram-golang/model/entity"
	"time"
)

type CreatePhotoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GetPhotoUser struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	Comments  string    `json:"comments"`
	// nambahin commenst dari mas Andri K
}

type GetPhotoDetailUser struct {
	ID        int         `json:"id"`
	Title     string      `json:"title"`
	Caption   string      `json:"caption"`
	PhotoURL  string      `json:"photo_url"`
	CreatedAt time.Time   `json:"created_at"`
	User      UserInPhoto `json:"user"`
	Comments  string      `json:"comments"`
	// nambahin commenst dari mas Andri K
}

type UserInPhoto struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GetAllPhotosUser(photos []entity.Photo) []GetPhotoUser {
	if len(photos) == 0 {
		return []GetPhotoUser{}
	}

	var allPhotoUser []GetPhotoUser

	for _, photo := range photos {
		tmpPhoto := GetPhotoUser{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			CreatedAt: photo.CreatedAt,
			Comments:  "Comments aja",
		}

		allPhotoUser = append(allPhotoUser, tmpPhoto)
	}

	return allPhotoUser
}
