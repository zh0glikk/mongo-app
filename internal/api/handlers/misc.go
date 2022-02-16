package handlers

import (
	"encoding/json"
	"github.com/google/jsonapi"
	"net/http"
)

func render(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}

func renderErr(w http.ResponseWriter, errs []*jsonapi.ErrorObject) {
	jsonapi.MarshalErrors(w, errs)
}

