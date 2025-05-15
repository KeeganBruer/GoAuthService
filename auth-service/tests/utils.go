package main_test

import (
	"bytes"
	"encoding/json"
	"kbrouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MakePostRequest(Router *kbrouter.Router, t *testing.T, path string, reqBody any) (*http.Response, error) {
	buf, err := MakeJSONBuffer(reqBody)
	if err != nil {
		t.Errorf("Could not make json into a buffer\n")
		return nil, err
	}
	//Construct request object and response recorder
	req := httptest.NewRequest(http.MethodPost, path, &buf)
	w := httptest.NewRecorder()

	//process the request through the app and produce a http response
	Router.ServeHTTP(w, req)
	httpRes := w.Result()
	return httpRes, nil
}

func MakeJSONBuffer(reqBody any) (bytes.Buffer, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(reqBody)
	return buf, err
}
