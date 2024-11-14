package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

// NewRestErr is a constructor for RestErr that accepts a message, error, HTTP
// status code, and optional causes.
func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// NewBadRequestError creates a new RestErr representing a bad request error
// with the specified message. It sets the error type to "bad request" and
// the HTTP status code to 400 (Bad Request).
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

// NewBadRequestValidationError creates a new RestErr representing a bad request
// error with the specified message and causes. It sets the error type to "bad
// request" and the HTTP status code to 400 (Bad Request). The causes parameter
// should be a slice of Causes, where each cause represents a field and its
// corresponding error message.
func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

// NewInternalServerError creates a new RestErr representing an internal server error
// with the specified message. It sets the error type to "internal server request"
// and the HTTP status code to 500 (Internal Server Error).
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal server request",
		Code:    http.StatusInternalServerError,
	}
}

// NewNotFoundError creates a new RestErr representing a not found error with
// the specified message. It sets the error type to "not found" and the HTTP
// status code to 404 (Not Found).
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "not found",
		Code:    http.StatusNotFound,
	}
}

// NewForbidenError creates a new RestErr representing a forbidden error
// with the specified message. It sets the error type to "forbidden" and
// the HTTP status code to 403 (Forbidden).
func NewForbidenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "forbiden",
		Code:    http.StatusForbidden,
	}
}
