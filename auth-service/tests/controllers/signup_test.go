package controller_tests

import (
	controller_signup "go-auth-service/controllers/signup"
	main_test "go-auth-service/tests"
	"testing"
)

func TestSignup(t *testing.T) {
	router := app.PublicRouter
	TargetUsername := "Test User"

	//Create Request Details
	reqBody := &controller_signup.SignupRequest{
		Username: TargetUsername,
	}
	httpRes, err := main_test.MakePostRequest(router, t, "/signup", reqBody)
	if err != nil {
		t.Errorf("Error sending request\n")
		return
	}

	if httpRes.StatusCode != 200 {
		t.Errorf("Status Code: %d\n", httpRes.StatusCode)
		return
	}

}
