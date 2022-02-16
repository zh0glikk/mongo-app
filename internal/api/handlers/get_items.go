package handlers

import (
	"net/http"

	"github.com/zh0glikk/mongo-app/internal/api/requests"
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetItemsRequest(r)
	if err != nil {
		renderErr(w, BadRequest(err))
		return
	}

	q := Items(r)

	if request.Title != nil {
		q = q.FilterByTitle(*request.Title)
	}
	if request.AmountLt != nil {
		q = q.FilterByAmountLt(*request.AmountLt)
	}
	if request.AmountGt != nil {
		q = q.FilterByAmountGt(*request.AmountGt)
	}

	items, err := q.GetBatch(r.Context())
	if err != nil {
		Log(r).WithError(err).Error("failed to get items")
		renderErr(w, InternalError())
		return
	}

	render(w, items)
}
