package utils_responses

import "net/http"

type RestMessage struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}


func CustomSuccessResponse(msg string) *RestMessage {
	return &RestMessage{
		Message: msg,
		Status:  http.StatusOK,
		Error:   "success_Request",
	}
}