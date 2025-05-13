package kbrouter

import (
	"encoding/json"
	"net/http"
)

type KBResponse struct {
	writer http.ResponseWriter
}

func (res *KBResponse) SendString(val string) {
	res.writer.Header().Set("Content-Type", "text/plain")
	res.writer.Write([]byte(val))
}
func (res *KBResponse) SendJSON(resval any) {
	res.writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res.writer).Encode(resval)
}
