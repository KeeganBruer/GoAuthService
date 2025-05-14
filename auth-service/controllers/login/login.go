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

	//Create a pair of JWT tokens with different expirations
	data := &jwttokens.NewTokenData{
		UserID: body.Username,
	}
	token, err := jwttokens.CreateToken(data)
	if err != nil {
		res.SendString("Could not create token")
		return
	}
	refreshToken, err := jwttokens.CreateToken(data)
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
