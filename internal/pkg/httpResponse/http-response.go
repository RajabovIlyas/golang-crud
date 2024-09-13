package httpResponse

import "net/http"

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

const (
	StatusOK        = "OK"
	StatusCreated   = "Created"
	StatusNoContent = "No Content"
)

func CreatedResponse(data interface{}) (int, Response) {
	return http.StatusCreated, Response{
		Code:   http.StatusCreated,
		Status: StatusCreated,
		Data:   data,
	}
}

func NoContentResponse(message string) (int, Response) {
	return http.StatusNoContent, Response{
		Code:   http.StatusNoContent,
		Status: StatusNoContent,
		Data:   MessageResponse{Message: message},
	}
}

func SuccessResponse(data interface{}) (int, Response) {
	return http.StatusOK, Response{
		Code:   http.StatusOK,
		Status: StatusOK,
		Data:   data,
	}
}
