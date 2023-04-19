package httpserver

import "net/http"

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	RootError  error  `json:"root_error"`
	ErrorLog   string `json:"error_log"`
	Message    string `json:"message"`
	ErrorCode  string `json:"error_code"`
}

func (e *ErrorResponse) RootErr() error {
	if err, ok := e.RootError.(*ErrorResponse); ok {
		return err.RootErr()
	}

	return e.RootError
}

func (e *ErrorResponse) Error() string {
	return e.RootErr().Error()
}

func FullErrorResponse(statusCode int, err error, errorLog, message, errorCode string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		RootError:  err,
		ErrorLog:   errorLog,
		Message:    message,
		ErrorCode:  errorCode,
	}
}

func SimpleErrorResponse(err error, message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusBadRequest,
		RootError:  err,
		ErrorLog:   err.Error(),
		Message:    message,
	}
}
