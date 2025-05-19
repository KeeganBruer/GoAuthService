package kbrouter

import (
	"encoding/json"
	"net/http"
)

type KBRequest struct {
	httpReq *http.Request
	Host    string
	Path    string
}

func (req *KBRequest) GetHeader(key string) []string {
	return req.httpReq.Header[key]
}

// Parse the content of a post request
func (req *KBRequest) ParseBodyJSON(out any) {
	decoder := json.NewDecoder(req.httpReq.Body)
	err := decoder.Decode(out)
	if err != nil {
		panic(err)
	}
}
