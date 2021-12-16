package input

type InputPhotos struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" binding:"required"`
}

type DeletePhoto struct {
	ID int `uri:"id" binding:"required"`
}
