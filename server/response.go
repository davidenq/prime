package server

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Message interface{} `json:"data"`
}

//Reply .
func Reply(w http.ResponseWriter, code int, data interface{}) {

	result := &response{
		Status:  "ok",
		Code:    http.StatusOK,
		Message: data,
	}
	w.Header().Set("Content-Type", "application/json")
	outcome, err := json.Marshal(result)

	if err != nil {
		result := &response{
			Status:  "error",
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		outcome, err = json.Marshal(result)
	}
	w.Write(outcome)
	return
}
