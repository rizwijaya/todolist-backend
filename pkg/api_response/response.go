package api

func NewErrorResponse(status string, message string) ResponseError {
	return ResponseError{
		Status:  status,
		Message: message,
	}
}

func NewSuccessResponse(status string, message string, data interface{}) ResponseSuccess {
	return ResponseSuccess{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
