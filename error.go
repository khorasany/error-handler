package error_handler

import "net/http"

type Error struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func WrongPassword() Error {
	return Error{
		ErrorCode: http.StatusForbidden,
		Message:   "wrong password,attention: limit attempt",
	}
}
