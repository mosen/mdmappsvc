package source

import (
	"bytes"
	"encoding/json"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
)

func MakeHTTPHandler(ctx context.Context, s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// GET		/sources/	get a list of all package sources
	// POST		/sources/	create a package source
	r.Methods("GET").Path("/v1/sources/").Handler(httptransport.NewServer(
		ctx,
		e.GetSourcesEndpoint,
		decodeGetSourcesRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/v1/sources/").Handler(httptransport.NewServer(
		ctx,
		e.PostSourceEndpoint,
		decodePostSourceRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodePostSourceRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req postSourceRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Source); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetSourcesRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	//var req getSourcesRequest
	//if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
	//	return nil, e
	//}
	//return req, nil
	return nil, nil
}

type errorer interface {
	error() error
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	default:
		if e, ok := err.(httptransport.Error); ok {
			switch e.Domain {
			case httptransport.DomainDecode:
				return http.StatusBadRequest
			case httptransport.DomainDo:
				return http.StatusServiceUnavailable
			default:
				return http.StatusInternalServerError
			}
		}
		return http.StatusInternalServerError
	}
}

//// CLIENT

func encodePostSourceRequest(ctx context.Context, req *http.Request, request interface{}) error {
	req.Method, req.URL.Path = "POST", "/v1/sources/"
	return encodeRequest(ctx, req, request)
}

func decodePostSourceResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response postSourceResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func encodeGetSourcesRequest(ctx context.Context, req *http.Request, request interface{}) error {
	req.Method, req.URL.Path = "GET", "/v1/sources/"
	return encodeRequest(ctx, req, request)
}

func decodeGetSourcesResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response getSourcesResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
