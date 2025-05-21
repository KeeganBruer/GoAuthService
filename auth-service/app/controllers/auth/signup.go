package auth_controller

import (
	"kbrouter"
)

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Post Request to the login endpoint
func (controller *AuthController) Signup_PostRequest(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
	var body SignupRequest
	req.ParseBodyJSON(&body)

	UserModel := controller.Models.GetUserModel()
	_, err := UserModel.GetUserByUsername(body.Username)
	if err == nil {
		res.SetStatusCode(409).SendString("User already exists")
		return
	}
	user := UserModel.NewUser()
	user.Username = body.Username
	user.SetPassword(body.Password)
	user.Save()

	res.SendString("OKAY")
}
