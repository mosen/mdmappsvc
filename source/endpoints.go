package source

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type Endpoints struct {
	PostSourceEndpoint   endpoint.Endpoint
	GetSourceEndpoint    endpoint.Endpoint
	PutSourceEndpoint    endpoint.Endpoint
	PatchSourceEndpoint  endpoint.Endpoint
	DeleteSourceEndpoint endpoint.Endpoint
}

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		PostSourceEndpoint: MakePostSourceEndpoint(s),
	}
}

func MakePostSourceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(postSourceRequest)
		e := s.PostSource(ctx, &req.Source)
		return postSourceResponse{Err: e}, nil
	}
}

type postSourceRequest struct {
	Source Source
}

type postSourceResponse struct {
	Err error `json:"err,omitempty"`
}