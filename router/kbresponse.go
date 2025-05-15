package kbrouter

import (
	"encoding/json"
	"net/http"
)

type KBResponse struct {
	writer http.ResponseWriter
}

// Send a string as the content of the response
func (res *KBResponse) SendString(val string) {
	res.writer.Header().Set("Content-Type", "text/plain")
	res.writer.Write([]byte(val))
}

// Send a JSON object as the content of the response
func (res *KBResponse) SendJSON(resval any) {
	res.writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res.writer).Encode(resval)
}

// Set the response status code
func (res *KBResponse) SetStatusCode(code int) *KBResponse {
	res.writer.WriteHeader(code)
	return res
}
