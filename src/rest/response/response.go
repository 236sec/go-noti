package response

import "net/http"

type BaseResponse[T any] struct {
	Success    bool   `json:"success"`
	Data       T      `json:"data,omitempty"`
	Message    string `json:"message,omitempty"`
	HttpStatus int    `json:"-"`
}

type ResponseCode int

const (
	SuccessResponse ResponseCode = iota + 1
	BadRequestResponse
	UnauthorizedResponse
	ForbiddenResponse
	NotFoundResponse
	ConflictResponse
	InternalServerErrorResponse
)

var Responses = map[ResponseCode]BaseResponse[any]{
	SuccessResponse:             {Success: true, Message: "Success", HttpStatus: http.StatusOK},
	BadRequestResponse:          {Success: false, Message: "Bad Request", HttpStatus: http.StatusBadRequest},
	UnauthorizedResponse:        {Success: false, Message: "Unauthorized", HttpStatus: http.StatusUnauthorized},
	ForbiddenResponse:           {Success: false, Message: "Forbidden", HttpStatus: http.StatusForbidden},
	NotFoundResponse:            {Success: false, Message: "Not Found", HttpStatus: http.StatusNotFound},
	ConflictResponse:            {Success: false, Message: "Conflict", HttpStatus: http.StatusConflict},
	InternalServerErrorResponse: {Success: false, Message: "Internal Server Error", HttpStatus: http.StatusInternalServerError},
}
