package source

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
	"strings"
	"net/url"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Endpoints struct {
	PostSourceEndpoint   endpoint.Endpoint
	GetSourceEndpoint    endpoint.Endpoint
	PutSourceEndpoint    endpoint.Endpoint
	PatchSourceEndpoint  endpoint.Endpoint
	DeleteSourceEndpoint endpoint.Endpoint
	GetSourcesEndpoint   endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		PostSourceEndpoint: MakePostSourceEndpoint(s),
		GetSourcesEndpoint: MakeGetSourcesEndpoint(s),
	}
}

func MakeClientEndpoints(instance string) (Endpoints, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}
	tgt.Path = ""

	options := []httptransport.ClientOption{}

	return Endpoints{
		PostSourceEndpoint:   httptransport.NewClient("POST", tgt, encodePostSourceRequest, decodePostSourceResponse, options...).Endpoint(),
		GetSourcesEndpoint:   httptransport.NewClient("GET", tgt, encodeGetSourcesRequest, decodeGetSourcesResponse, options...).Endpoint(),
	}, nil
}

//// CLIENT

func (e Endpoints) PostSource(ctx context.Context, s *Source) error {
	request := postSourceRequest{Source: *s}
	response, err := e.PostSourceEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(postSourceResponse)
	return resp.Err
}

func (e Endpoints) GetSource(ctx context.Context, uuidStr string) (Source, error) {
	request := getSourceRequest{UUID: uuidStr}
	response, err := e.GetSourceEndpoint(ctx, request)
	if err != nil {
		return Source{}, err
	}
	resp := response.(getSourceResponse)
	return resp.Source, resp.Err
}

func (e Endpoints) PutSource(ctx context.Context, uuidStr string, s *Source) error {
	request := putSourceRequest{UUID: uuidStr, Source: *s}
	response, err := e.PutSourceEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(putSourceResponse)
	return resp.Err
}

func (e Endpoints) PatchSource(ctx context.Context, uuidStr string, s *Source) error {
	request := patchSourceRequest{UUID: uuidStr, Source: *s}
	response, err := e.PatchSourceEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(patchSourceResponse)
	return resp.Err
}

func (e Endpoints) DeleteSource(ctx context.Context, uuidStr string) error {
	request := deleteSourceRequest{UUID: uuidStr}
	response, err := e.DeleteSourceEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(deleteSourceResponse)
	return resp.Err
}

func (e Endpoints) GetSources(ctx context.Context) ([]Source, error) {
	request := getSourcesRequest{}
	response, err := e.GetSourcesEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	resp := response.(getSourcesResponse)
	return resp.Sources, resp.Err
}



//// SERVER

func MakePostSourceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postSourceRequest)
		e := s.PostSource(ctx, &req.Source)
		return postSourceResponse{Err: e}, nil
	}
}

func MakeGetSourcesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(getSourcesRequest)
		sources, e := s.GetSources(ctx)
		return getSourcesResponse{Err: e, Sources: sources}, nil
	}
}

type postSourceRequest struct {
	Source Source
}

type postSourceResponse struct {
	Err error `json:"err,omitempty"`
}

func (r postSourceResponse) error() error { return r.Err }

type getSourceRequest struct {
	UUID string
}

type getSourceResponse struct {
	Source Source `json:"source,omitempty"`
	Err error `json:"err,omitempty"`
}

func (r getSourceResponse) error() error { return r.Err }

type putSourceRequest struct {
	UUID string
	Source Source
}

type putSourceResponse struct {
	Err error `json:"err,omitempty"`
}

func (r putSourceResponse) error() error { return r.Err }

type patchSourceRequest struct {
	UUID string
	Source Source
}

type patchSourceResponse struct {
	Err error `json:"err,omitempty"`
}

func (r patchSourceResponse) error() error { return r.Err }

type deleteSourceRequest struct {
	UUID string
}

type deleteSourceResponse struct {
	Err error `json:"err,omitempty"`
}

func (r deleteSourceResponse) error() error { return r.Err }

type getSourcesRequest struct {}

type getSourcesResponse struct {
	Sources []Source `json:"sources,omitempty"`
	Err error `json:"err,omitempty"`
}

func (r getSourcesResponse) error() error { return r.Err }
