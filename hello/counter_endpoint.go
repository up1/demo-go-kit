package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type addRequest struct {
	V int `json:"value"`
}

type addResponse struct {
	V int `json:"value"`
}

func makeAddEndpoint(sv Counter) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		v := sv.Add(req.V)
		return addResponse{v}, nil
	}
}

func decodeAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req addRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
