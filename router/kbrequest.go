package kbrouter

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type KBRequest struct {
	httpReq    *http.Request
	Host       string
	CurrPath   string
	Path       string
	Parameters map[string]string
}

func NewKBRequest(httpReq *http.Request, basepath string) *KBRequest {
	CurrPath := strings.Replace(httpReq.URL.Path, basepath, "", 1)
	//Empty route should map to the / route
	if CurrPath == "" {
		CurrPath = "/"
	}
	req := &KBRequest{
		httpReq:    httpReq,
		Host:       httpReq.URL.Host,
		CurrPath:   CurrPath,
		Path:       httpReq.URL.Path,
		Parameters: make(map[string]string),
	}
	return req
}
func (req *KBRequest) GetHeader(key string) []string {
	return req.httpReq.Header[key]
}
func (req *KBRequest) GetParam(key string) string {
	return req.Parameters[key]
}
func (req *KBRequest) GetIntParam(key string) (int, error) {
	str := req.Parameters[key]
	return strconv.Atoi(str)
}

// Parse the content of a post request
func (req *KBRequest) ParseBodyJSON(out any) {
	decoder := json.NewDecoder(req.httpReq.Body)
	err := decoder.Decode(out)
	if err != nil {
		panic(err)
	}
}
