package utils

// ErrorPayload to be type cast into the appropriate Error and returned by go-swagger application
type ErrorPayload struct {
	Code    int64
	Message *string
}

// GetErrorPayload creates the error payload to return
func GetErrorPayload(code int, err error) ErrorPayload {
	message := err.Error()
	return ErrorPayload{
		Code:    int64(code),
		Message: &message,
	}
}
