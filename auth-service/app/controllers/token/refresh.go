package token_controller

import (
	"go-auth-service/app/services/jwttokens"
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
func (controller *TokenController) Refresh_PostRequest(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
	var body RefreshRequest
	req.ParseBodyJSON(&body)

	existing_token, err := jwttokens.DecodeToken(body.RefreshToken)
	if err != nil {
		res.SetStatusCode(400).SendString("Error decoding token")
		return
	}
	if existing_token.Type != "refresh_token" {
		res.SetStatusCode(400).SendString("Token is not a session refresh token")
		return
	}

	SessionModel := controller.Models.GetSessionModel()
	session, err := SessionModel.GetSessionByRefreshID(existing_token.ID)
	if err != nil {
		res.SetStatusCode(400).SendString("Could not find session")
		return
	}
	tokens, err := session.GetTokens()
	if err != nil {
		res.SetStatusCode(400).SendString("Could not get session tokens")
		return
	}

	//Construct and send response
	resVal := &RefreshResponse{
		Token:   tokens.Token,
		Refresh: tokens.Refresh,
	}
	res.SendJSON(resVal)
}
