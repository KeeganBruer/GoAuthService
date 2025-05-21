package kbrouter

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type KBResponse struct {
	writer http.ResponseWriter
	IsOpen bool
}

func NewKBResponse(w http.ResponseWriter) *KBResponse {
	res := &KBResponse{
		writer: w,
		IsOpen: true,
	}
	return res
}

func (res *KBResponse) Close() {
	res.IsOpen = false
}
func (res *KBResponse) SetHeader(key string, val string) *KBResponse {
	if !res.IsOpen {
		return res
	}
	res.writer.Header().Set(key, val)
	return res
}

// Send a string as the content of the response
func (res *KBResponse) SendString(val string) error {
	if !res.IsOpen {
		return errors.New("response is closed")
	}
	if res.writer.Header().Get("Content-Type") == "" {
		res.SetHeader("Content-Type", "text/plain")
	}
	res.writer.Write([]byte(val))
	return nil
}

// Send a JSON object as the content of the response
func (res *KBResponse) SendJSON(resval any) error {
	if !res.IsOpen {
		return errors.New("response is closed")
	}
	if res.writer.Header().Get("Content-Type") == "" {
		res.SetHeader("Content-Type", "application/json")
	}
	json.NewEncoder(res.writer).Encode(resval)
	return nil
}
func (res *KBResponse) SendJSONString(fmtStr string, vals ...any) error {
	if !res.IsOpen {
		return errors.New("response is closed")
	}
	if res.writer.Header().Get("Content-Type") == "" {
		res.SetHeader("Content-Type", "application/json")
	}
	data := fmt.Sprintf(fmtStr, vals...)
	res.writer.Write([]byte(data))
	return nil
}
func (res *KBResponse) SendFile(filepath string) error {
	if !res.IsOpen {
		return errors.New("response is closed")
	}
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	if res.writer.Header().Get("Content-Type") == "" {
		mimeType := http.DetectContentType(data)
		res.SetHeader("Content-Type", mimeType)
	}
	res.writer.Write(data)
	return nil
}

// Set the response status code
func (res *KBResponse) SetStatusCode(code int) *KBResponse {
	if !res.IsOpen {
		return res
	}
	res.writer.WriteHeader(code)
	return res
}
