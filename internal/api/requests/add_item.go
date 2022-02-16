package requests

import (
	"encoding/json"
	"net/http"
)

type CreateItemRequest struct {
	Title  string `json:"name"`
	Amount uint64 `json:"amount"`
	Price  uint64 `json:"price"`
}

func NewCreateItemRequest(r *http.Request) (*CreateItemRequest, error) {
	var request CreateItemRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	return &request, nil
}
