package httpserver

import "net/http"

type SuccessResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func FullSuccessResponse(statusCode int, message string, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}

func SimpleSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       data,
	}
}
