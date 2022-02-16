package requests

import (
	"net/http"
	"strconv"
)

type GetItemsRequest struct {
	AmountGt *uint64
	AmountLt *uint64
	Title    *string
}

func NewGetItemsRequest(r *http.Request) (*GetItemsRequest, error) {
	var req GetItemsRequest

	values := r.URL.Query()

	amountGt, ok := values["amount_gt"]
	if ok {
		_amountGt, err := strconv.ParseInt(amountGt[0], 0, 64)
		if err != nil {
			return nil, err
		}
		q := uint64(_amountGt)

		req.AmountGt = &(q)
	}
	amountLt, ok := values["amount_lt"]
	if ok {
		_amountLt, err := strconv.ParseInt(amountLt[0], 0, 64)
		if err != nil {
			return nil, err
		}
		q := uint64(_amountLt)

		req.AmountLt = &(q)
	}

	title, ok := values["title"]
	if ok {
		req.Title = &title[0]
	}

	return &req, nil
}
