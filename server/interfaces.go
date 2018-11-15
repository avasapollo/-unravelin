package server

type RestApiServer interface {
	ListenServe(port int)
}

type Validation interface {
	ValidateFormRequest(request map[string]interface{}) error
}
