package controller_token

import (
	"go-auth-service/services/jwttokens"
	"kbrouter"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}
type RefreshResponse struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
}

// Post Request to the login endpoint
func Refresh_PostRequest(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
	var body RefreshRequest
	req.ParseBodyJSON(&body)

	existing_token, err := jwttokens.DecodeToken(body.RefreshToken)
	if err != nil {
		res.SetStatusCode(400).SendString("Error decoding token")
		return
	}
	//Create a pair of JWT tokens with different expirations
	token, err := jwttokens.CreateToken(&jwttokens.NewTokenData{
		UserID:        existing_token.UserID,
		MinutesTilExp: 30,
	})
	if err != nil {
		res.SetStatusCode(400).SendString("Could not create token")
		return
	}
	refreshToken, err := jwttokens.CreateToken(&jwttokens.NewTokenData{
		UserID:        existing_token.UserID,
		MinutesTilExp: 60,
	})
	if err != nil {
		res.SetStatusCode(400).SendString("Could not create refresh token")
		return
	}
	//Construct and send response
	resVal := &RefreshResponse{
		Token:   token,
		Refresh: refreshToken,
	}
	res.SendJSON(resVal)
}
