package response

import (
	"encoding/json"
	"net/http"
)

const (
	unmarshallingError = `{"error": "an error occurred during the response please try again"}`
)

type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func String(str string) *string {
	return &str
}

func Write(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	msg, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(unmarshallingError))
		return
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
