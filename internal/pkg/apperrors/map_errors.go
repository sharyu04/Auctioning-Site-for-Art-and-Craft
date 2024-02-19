package apperrors

import "net/http"

type res struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_msg"`
}

func MapError(err error) res {
	switch err.(type) {
	case BadRequest:
		return res{ErrorCode: http.StatusBadRequest, ErrorMessage: err.Error()}
	case NoContent:
		return res{ErrorCode: http.StatusNotFound, ErrorMessage: err.Error()}
	case UnAuthorizedAccess:
		return res{ErrorCode: http.StatusUnauthorized, ErrorMessage: err.Error()}
	default:
		return res{ErrorCode: http.StatusInternalServerError, ErrorMessage: err.Error()}
	}

}
