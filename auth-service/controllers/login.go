package controllers

import (
	"fmt"
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
func Request_Post_Login(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
	fmt.Println("got POST request to /login")

	var body LoginRequest
	req.ParseBodyJSON(&body)
	data := &jwttokens.NewTokenData{
		UserID: body.Username,
	}
	token, err := jwttokens.CreateToken(data)
	if err != nil {
		res.SendString("ERROR")
		return
	}
	refreshToken, err := jwttokens.CreateToken(data)
	if err != nil {
		res.SendString("ERROR")
		return
	}

	resVal := &LoginResponse{
		Token:   token,
		Refresh: refreshToken,
	}
	res.SendJSON(resVal)
}
