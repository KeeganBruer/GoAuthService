package controller_signup

import (
	"kbrouter"
)

type SignupRequest struct {
	Username string `json:"username"`
}

// Post Request to the login endpoint
func Signup_PostRequest(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
	var body SignupRequest
	req.ParseBodyJSON(&body)

	res.SendString("OKAY")
}
