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

func HttpErrorResponse(err error) *ErrorResponse {
	if e, ok := err.(*ErrorResponse); ok {
		return e
	} else {
		return &ErrorResponse{
			StatusCode: http.StatusBadRequest,
			RootError:  err,
			ErrorLog:   logFromError(err),
			Message:    "unknown error",
			ErrorCode:  "UNKNOWN_ERROR",
		}
	}
}

func logFromError(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
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

func BadRequestErrorResponse(err error, message string, errorCode string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusBadRequest,
		RootError:  err,
		ErrorLog:   logFromError(err),
		Message:    message,
		ErrorCode:  errorCode,
	}
}

func InternalErrorResponse(err error, message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusInternalServerError,
		RootError:  err,
		ErrorLog:   logFromError(err),
		Message:    message,
		ErrorCode:  "INTERNAL_SERVER_ERROR",
	}
}

func AuthRequireErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusNetworkAuthenticationRequired,
		RootError:  err,
		ErrorLog:   err.Error(),
		Message:    "",
		ErrorCode:  "AUTHENTICATION_REQUIRE",
	}
}
