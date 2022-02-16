package handlers

import (
	"fmt"
	"github.com/google/jsonapi"
	"net/http"
)

func BadRequest(err error) []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		Title:  http.StatusText(http.StatusBadRequest),
		Status: fmt.Sprintf("%d", http.StatusBadRequest),
		Detail: "Your request was invalid in some way",
		Meta: &map[string]interface{}{
			"error": err.Error(),
		},
	}}
}

func NotFound() []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		Title:  http.StatusText(http.StatusNotFound),
		Status: fmt.Sprintf("%d", http.StatusNotFound),
	}}
}

func InternalError() []*jsonapi.ErrorObject {
	return []*jsonapi.ErrorObject{{
		Title:  http.StatusText(http.StatusInternalServerError),
		Status: fmt.Sprintf("%d", http.StatusInternalServerError),
	}}
}
