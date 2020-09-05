package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/netgame/backend/internal/model"
)

const serverErrorMessage = `{ "message": "Server error" }`

func getBodyContent(r *http.Request, data interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, r.ContentLength))
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

func returnError(w http.ResponseWriter, out model.Output) {
	w.WriteHeader(out.Status)
	if err := json.NewEncoder(w).Encode(out); err != nil {
		http.Error(w, serverErrorMessage, http.StatusInternalServerError)
		return
	}
}

func returnJson(w http.ResponseWriter, v interface{}) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, serverErrorMessage, http.StatusInternalServerError)
		return
	}
}
