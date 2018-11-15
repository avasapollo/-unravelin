package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/avasapollo/unravelin/data"
	"github.com/avasapollo/unravelin/encoder"
	"github.com/avasapollo/unravelin/printer"
	"github.com/gorilla/mux"
)

type ApiRest struct {
	printer    printer.Printer
	parser     data.Parser
	router     *mux.Router
	encoder    encoder.HashEncoder
	validation Validation
}

func NewApiRest(printer printer.Printer, parser data.Parser, encoder encoder.HashEncoder) RestApiServer {
	api := &ApiRest{
		printer:    printer,
		parser:     parser,
		router:     nil,
		encoder:    encoder,
		validation: NewValidation(),
	}

	rr := mux.NewRouter()

	// add handles
	// health return 200
	rr.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {}).Methods(http.MethodGet)
	rr.Handle("/v1/form",
		api.middleWareJsonFormat(http.HandlerFunc(api.PostForm))).Methods(http.MethodPost)

	api.router = rr
	return api
}

func (api ApiRest) PostForm(w http.ResponseWriter, r *http.Request) {
	request := make(map[string]interface{})
	// parse request
	if err := api.parseHttpRequest(r, &request); err != nil {
		api.write(w, http.StatusBadRequest,
			NewErrorResponse(http.StatusBadRequest, "couldn't possible parse the body"))
		return
	}

	// validation
	if err := api.validation.ValidateFormRequest(request); err != nil {
		api.write(w, http.StatusBadRequest,
			NewErrorResponse(http.StatusBadRequest, "couldn't possible parse the body"))
		return
	}

	// first print request
	api.printer.Print("http request (map[string]interface{})", request)

	// parse to data
	res, err := api.parser.ParseMapToData(request)
	if err != nil {
		api.write(w, http.StatusBadRequest,
			NewErrorResponse(http.StatusBadRequest, "couldn't possible deserialize the body"))
		return
	}

	// second print data structure
	api.printer.Print("data structure after the parsing (Data)", res)

	// get hash
	h := api.encoder.Encode(res.WebsiteUrl)

	// third print hash
	api.printer.Print("the hash is (string) ", h)

	// build response
	response := api.buildResponse(res, h)

	// forth print is the response
	api.printer.Print("the http response", response)

	// return the data structure to frontend
	api.write(w, http.StatusAccepted, response)
}

func (api ApiRest) buildResponse(data *data.Data, hash string) DataResponse {
	return DataResponse{
		WebsiteUrl: data.WebsiteUrl,
		SessionId:  data.SessionId,
		ResizeFrom: struct {
			Width  string
			Height string
		}{
			Width:  data.ResizeFrom.Width,
			Height: data.ResizeFrom.Height,
		},
		ResizeTo: struct {
			Width  string
			Height string
		}{
			Width:  data.ResizeTo.Width,
			Height: data.ResizeTo.Height,
		},
		CopyAndPaste:       data.CopyAndPaste,
		FormCompletionTime: data.FormCompletionTime,
		Hash:               hash,
	}
}

func (api ApiRest) middleWareJsonFormat(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-type")
		if contentType == "" {
			api.write(w, http.StatusBadRequest,
				NewErrorResponse(http.StatusBadRequest,
					"you need to specify content-type=application/json"))
			return
		}

		if contentType != "application/json" {
			api.write(w, http.StatusBadRequest,
				NewErrorResponse(http.StatusBadRequest,
					fmt.Sprintf("%s is not allowed, you have to use application/json")))
			return
		}
		// application/json content
		h.ServeHTTP(w, r)
	})
}

func (api ApiRest) ListenServe(port int) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), api.router))
}

func (api ApiRest) write(w http.ResponseWriter, code int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (api ApiRest) parseHttpRequest(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&target)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}
