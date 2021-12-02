package helper

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func APIResponse(status string, data interface{}) Response {
	jsonResponse := Response{
		Status: status,
		Data:   data,
	}

	return jsonResponse
}
