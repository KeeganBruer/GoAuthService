package controller_login

import (
	"go-auth-service/models"
	"kbrouter"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
}

// Post Request to the login endpoint
func Login_PostRequest(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
	var body LoginRequest
	req.ParseBodyJSON(&body)

	user, err := models.GetUserByUsername(body.Username)
	if err != nil {
		res.SetStatusCode(404).SendString("Could not find user")
		return
	}
	if !user.CheckPassword(body.Password) {
		res.SetStatusCode(401).SendString("Incorrect password")
		return
	}
	session := models.CreateOrGetSession(&models.NewSession{
		UserID: user.ID,
	})

	tokens, err := session.GetTokens()
	if err != nil {
		res.SetStatusCode(400).SendString("Could not get session tokens")
		return
	}
	//Construct and send response
	resVal := &LoginResponse{
		Token:   tokens.Token,
		Refresh: tokens.Refresh,
	}
	res.SendJSON(resVal)
}
