package handlers

import (
	"net/http"

	"github.com/zh0glikk/mongo-app/internal/api/requests"
	"github.com/zh0glikk/mongo-app/internal/data"
)

func AddItem(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateItemRequest(r)
	if err != nil {
		renderErr(w, BadRequest(err))
		return
	}

	_, err = Items(r).Create(r.Context(), data.Item{
		Title:  request.Title,
		Amount: request.Amount,
		Price:  request.Price,
	})
	if err != nil {
		Log(r).WithError(err).Error("failed to create item")
		renderErr(w, InternalError())
		return
	}

	render(w, http.StatusCreated)
}
