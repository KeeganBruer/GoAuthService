package controller_tests

import (
	"bytes"
	"encoding/json"
	controller_login "go-auth-service/controllers/login"
	"go-auth-service/services/jwttokens"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {

	TargetUsername := "Test User"

	//Create Request Details
	reqBody := &controller_login.LoginRequest{
		Username: TargetUsername,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(reqBody)
	if err != nil {
		log.Fatal(err)
	}

	//Construct request object and response recorder
	req := httptest.NewRequest(http.MethodPost, "/login", &buf)
	w := httptest.NewRecorder()

	//process the request through the app and produce a http response
	router.ServeHTTP(w, req)
	httpRes := w.Result()

	if httpRes.StatusCode != 200 {
		t.Errorf("Status Code: %d\n", httpRes.StatusCode)
		return
	}

	//Convert http reponse to Login response object
	var res controller_login.LoginResponse
	json.NewDecoder(httpRes.Body).Decode(&res)

	if res.Token == "" {
		t.Error("Empty Token\n")
		return
	}
	if res.Refresh == "" {
		t.Error("Empty Refresh Token\n")
		return
	}

	data, err := jwttokens.DecodeToken(res.Token)
	if err != nil {
		t.Error("Could not decode jwt token\n")
		return
	}
	if data.UserID != TargetUsername {
		t.Error("UserID does not match the TargetUsername\n")
		return
	}
}
