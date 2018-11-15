package server

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
	}
}

type DataResponse struct {
	WebsiteUrl string
	SessionId  string
	ResizeFrom struct {
		Width  string
		Height string
	}
	ResizeTo struct {
		Width  string
		Height string
	}
	CopyAndPaste       map[string]bool
	FormCompletionTime int
	Hash               string
}
