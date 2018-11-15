package server

import "github.com/gorilla/mux"

type RestApiServer interface {
	ListenServe(port int)
	GetMuxRouter() *mux.Router
}

type Validation interface {
	ValidateFormRequest(request map[string]interface{}) error
}
