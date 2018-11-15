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
	WebsiteUrl string `json:"websiteUrl"`
	SessionId  string `json:"sessionId"`
	ResizeFrom struct {
		Width  string `json:"width"`
		Height string `json:"height"`
	} `json:"resizeFrom"`
	ResizeTo struct {
		Width  string `json:"width"`
		Height string `json:"height"`
	} `json:"resizeTo"`
	CopyAndPaste       map[string]bool `json:"copyAndPaste"`
	FormCompletionTime int             `json:"formCompletionTime"`
	Hash               string          `json:"hash"`
}
