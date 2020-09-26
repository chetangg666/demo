package app

import (
	"app/metrics"
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints are exposed
type Endpoints struct {
	GetEndpoint    endpoint.Endpoint
	CreateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
}

// MakeGetEndpoint returns the response from our service "get"
func MakeGetEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(getRequest) // we really just need the request, we don't use any value from it
		d, err := srv.Get(ctx)
		if err != nil {
			return getResponse{d, err.Error()}, nil
		}
		return getResponse{d, ""}, nil
	}
}

// MakeCreateEndpoint returns the response from our service "ID"
func MakeCreateEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest) // we really just need the request, we use value from it
		s, err := srv.Create(ctx, req)
		if err != nil {
			metrics.RequestsToIndex.Inc()
			return addResponse{s}, err
		}
		metrics.RequestsToIndex.Inc()
		return addResponse{s}, nil
	}
}

// MakeDeleteEndpoint returns the response from our service "ID"
func MakeDeleteEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		b, err := srv.Delete(ctx, req.ID)
		if err != nil {
			metrics.RequestsToIndex.Inc()
			return deleteResponse{b, err.Error()}, nil
		}
		metrics.RequestsToIndex.Inc()
		return deleteResponse{b, ""}, nil
	}
}
