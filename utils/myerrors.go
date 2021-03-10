package utils

import (
	"errors"
	"net/http"
)

type YRestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func YCustomError(msg string) error{
	return errors.New(msg)
}

func YCustomBadRequestError(msg string) *YRestErr {
	return &YRestErr{
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   "Bad_Request",
	}
}

func YCustomNotFoundError(msg string) *YRestErr {
	return &YRestErr{
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func YCustomInternalServerError(msg string) *YRestErr {
	return &YRestErr{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}