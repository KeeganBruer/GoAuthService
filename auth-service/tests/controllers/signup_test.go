package controller_tests

import (
	"bytes"
	"encoding/json"
	controller_signup "go-auth-service/controllers/signup"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignup(t *testing.T) {

	TargetUsername := "Test User"

	//Create Request Details
	reqBody := &controller_signup.SignupRequest{
		Username: TargetUsername,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(reqBody)
	if err != nil {
		log.Fatal(err)
		return
	}

	//Construct request object and response recorder
	req := httptest.NewRequest(http.MethodPost, "/signup", &buf)
	w := httptest.NewRecorder()

	//process the request through the app and produce a http response
	router.ServeHTTP(w, req)
	httpRes := w.Result()

	if httpRes.StatusCode != 200 {
		t.Errorf("Status Code: %d\n", httpRes.StatusCode)
		return
	}

}
