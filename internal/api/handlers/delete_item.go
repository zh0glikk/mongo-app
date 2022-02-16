package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	item, err := Items(r).Delete(r.Context(), id)
	if err != nil {
		Log(r).WithError(err).Error("failed to get item")
		renderErr(w, InternalError())
		return
	}
	if item == nil {
		renderErr(w, NotFound())
		return
	}

	render(w, http.StatusNoContent)
}
