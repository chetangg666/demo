package app

import (
	"context"
	"encoding/json"
	"app/metrics"
	"net/http"
)

// In the first part of the file we are mapping requests and responses to their JSON payload.
type getRequest struct{}

type getResponse struct {
	Date string `json:"date"`
	Err  string `json:"err,omitempty"`
}

type deleteRequest struct {
	Date string `json:"date"`
}

type deleteResponse struct {
	ID  int    `json:="ID"`
	Err string `json:"err,omitempty"`
}

type addRequest struct {
	ID        int    `json:="ID"`
	LastName  string `json:"LastName"`
	FirstName string `json:"FirstName"`
	MsgType   string `json:"MsgType"`
}

type addResponse struct {
	ID int `json:="ID"`
}

// In the second part we will write "decoders" for our incoming requests
func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getRequest
	return req, nil
}

func decodeDeleteRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req addRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		metrics.RequestsToIndex.Inc()
		return nil, err
	}
	return req, nil
}

func decodeAddRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req addRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		metrics.RequestsToIndex.Inc()
		return nil, err
	}
	return req, nil
}

// Last but not least, we have the encoder for the response output
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
