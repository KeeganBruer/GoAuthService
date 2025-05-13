package controller_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	app "go-auth-service"
	"go-auth-service/controllers"
	"go-auth-service/services/jwttokens"
	"kbrouter"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *kbrouter.Router

func TestMain(m *testing.M) {
	router = app.CreateApp()
	m.Run()
}

func TestLogin(t *testing.T) {
	fmt.Println("Running Login Test")

	//Create Request Details
	reqBody := &controllers.LoginRequest{
		Username: "Test User",
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
	var res controllers.LoginResponse
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

	fmt.Println("data", data.UserID)
	fmt.Println("Successfully completed login test")
}
