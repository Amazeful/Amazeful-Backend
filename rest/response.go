package rest

type SuccessResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
}

func NewSuccessResponse(data interface{}, status int, message string) *SuccessResponse {
	return &SuccessResponse{
		Data:    data,
		Status:  status,
		Message: message,
	}
}

type ErrorResponse struct {
	Data   interface{} `json:"data,omitempty"`
	Status int         `json:"status"`
	Error  string      `json:"error,omitempty"`

	err error
}

func NewErrorResponse(data interface{}, status int, message string, err error) *ErrorResponse {
	return &ErrorResponse{
		Data:   data,
		Status: status,
		Error:  message,
		err:    err,
	}
}
