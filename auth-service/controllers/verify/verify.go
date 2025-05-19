package controller_verify

import (
	"go-auth-service/services/jwttokens"
	"kbrouter"
	"strings"
)

type VerifyResponse struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
}

// Post Request to the login endpoint
func Verify_GetRequest(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
	authorization := req.GetHeader("Authorization")[0]
	authorization = strings.Replace(authorization, "Bearer ", "", 1)
	token, err := jwttokens.DecodeToken(authorization)
	if err != nil {
		res.SendString("Error decoding token")
		return
	}
	res.SendJSON(token)
}
