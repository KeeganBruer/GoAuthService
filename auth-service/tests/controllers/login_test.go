package controller_tests

import (
	"encoding/json"
	"fmt"
	controller_login "go-auth-service/controllers/login"
	"go-auth-service/services/jwttokens"
	main_test "go-auth-service/tests"
	"testing"
)

func TestLogin(t *testing.T) {
	router := app.PublicRouter
	TargetUsername := "Test User"

	//Create Request Details
	reqBody := &controller_login.LoginRequest{
		Username: TargetUsername,
	}
	httpRes, err := main_test.MakePostRequest(router, t, "/login", reqBody)
	if err != nil {
		t.Errorf("Error sending request\n")
		return
	}
	if httpRes.StatusCode != 200 {
		t.Errorf("Status Code: %d\n", httpRes.StatusCode)
		return
	}

	//Convert http reponse to Login response object
	var res controller_login.LoginResponse
	json.NewDecoder(httpRes.Body).Decode(&res)

	// Confirm tokens are not empty
	if res.Token == "" {
		t.Error("Empty Token\n")
		return
	} else {
		t.Logf("Token Not Empty\n")
	}
	if res.Refresh == "" {
		t.Error("Empty Refresh Token\n")
		return
	} else {
		t.Logf("Refresh Token Not Empty\n")
	}

	//Check if token can be decoded
	data, err := jwttokens.DecodeToken(res.Token)
	if err != nil {
		t.Error("Could not decode jwt token\n")
		return
	}
	fmt.Printf("%v", data)
}
