package utils

// Payload struct to represent the error to return
type Payload struct {
	Code    int64
	Message *string
}

// GetErrorPayload creates the error payload to return
func GetErrorPayload(code int, err error) *Payload {
	message := err.Error()
	return &Payload{
		Code:    int64(code),
		Message: &message,
	}
}
