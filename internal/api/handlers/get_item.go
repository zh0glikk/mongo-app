package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	item, err := Items(r).Get(r.Context(), id)
	if err != nil {
		Log(r).WithError(err).Error("failed to get item")
		renderErr(w, InternalError())
		return
	}
	if item == nil {
		renderErr(w, NotFound())
		return
	}

	render(w, item)
}
