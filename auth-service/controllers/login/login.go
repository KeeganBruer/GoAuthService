package controller_login

import (
	"go-auth-service/services/jwttokens"
	"kbrouter"
)

type LoginRequest struct {
	Username string `json:"username"`
}
type LoginResponse struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
}

// Post Request to the login endpoint
func Login_PostRequest(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
	var body LoginRequest
	req.ParseBodyJSON(&body)

	//user_model := models.GetUserModel()

	//Create a pair of JWT tokens with different expirations
	token, err := jwttokens.CreateToken(&jwttokens.NewTokenData{
		UserID:        body.Username,
		MinutesTilExp: 30,
	})
	if err != nil {
		res.SendString("Could not create token")
		return
	}
	refreshToken, err := jwttokens.CreateToken(&jwttokens.NewTokenData{
		UserID:        body.Username,
		MinutesTilExp: 60,
	})
	if err != nil {
		res.SendString("Could not create refresh token")
		return
	}

	//Construct and send response
	resVal := &LoginResponse{
		Token:   token,
		Refresh: refreshToken,
	}
	res.SendJSON(resVal)
}
