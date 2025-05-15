package controller_signup

import (
	"go-auth-service/models"
	"kbrouter"
)

type SignupRequest struct {
	Username string `json:"username"`
}

// Post Request to the login endpoint
func Signup_PostRequest(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
	var body SignupRequest
	req.ParseBodyJSON(&body)

	user := models.NewUser()
	user.Username = body.Username
	user.Save()

	res.SendString("OKAY")
}
