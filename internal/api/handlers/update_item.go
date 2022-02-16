package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	mongo2 "go.mongodb.org/mongo-driver/mongo"

	"github.com/zh0glikk/mongo-app/internal/api/requests"
	"github.com/zh0glikk/mongo-app/internal/data"
)

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateItemRequest(r)
	if err != nil {
		renderErr(w, BadRequest(err))
		return
	}
	id := chi.URLParam(r, "id")

	err = Items(r).Update(r.Context(), id, data.Item{
		Title:  request.Title,
		Amount: request.Amount,
		Price:  request.Price,
	})
	if err != nil {
		if err == mongo2.ErrNoDocuments {
			renderErr(w, NotFound())
			return
		}

		Log(r).WithError(err).Error("failed to update item")
		renderErr(w, InternalError())
		return
	}

	render(w, http.StatusNoContent)
}
