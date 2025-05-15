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

	user_model := models.GetUserModel()
	user_model.AddUser(&models.User{
		Username: body.Username,
	})

	res.SendString("OKAY")
}
