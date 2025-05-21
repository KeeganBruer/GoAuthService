package token_controller

import (
	"go-auth-service/app/services/jwttokens"
	"kbrouter"
	"strings"
)

type VerifyResponse struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
}

// Post Request to the login endpoint
func (controller *TokenController) Verify_GetRequest(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
	authorization := req.GetHeader("Authorization")[0]
	authorization = strings.Replace(authorization, "Bearer ", "", 1)
	token, err := jwttokens.DecodeToken(authorization)
	if err != nil {
		res.SetStatusCode(400).SendString("Error decoding token " + err.Error())
		return
	}
	// if token.Type != "session" {
	// 	res.SetStatusCode(400).SendString("Not a session token")
	// 	return
	// }
	res.SendJSON(token)
}
