package kbrouter

import (
	"encoding/json"
	"net/http"
	"os"
)

type KBResponse struct {
	writer http.ResponseWriter
}

func (res *KBResponse) SetHeader(key string, val string) {
	res.writer.Header().Set(key, val)
}

// Send a string as the content of the response
func (res *KBResponse) SendString(val string) {
	res.SetHeader("Content-Type", "text/plain")
	res.writer.Write([]byte(val))
}

// Send a JSON object as the content of the response
func (res *KBResponse) SendJSON(resval any) {
	res.SetHeader("Content-Type", "application/json")
	json.NewEncoder(res.writer).Encode(resval)
}
func (res *KBResponse) SendFile(filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	res.writer.Write(data)
	//fmt.Println(string(data))
	return nil
}

// Set the response status code
func (res *KBResponse) SetStatusCode(code int) *KBResponse {
	res.writer.WriteHeader(code)
	return res
}
