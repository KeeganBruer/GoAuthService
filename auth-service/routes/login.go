package routes

import (
	"fmt"
	"go-auth-service/jwttokens"
	"kbrouter"
)

type LoginRequest struct {
	Username string `json:"username"`
}
type LoginResponse struct {
	Test string `json:"test"`
}

// Post Request to the login endpoint
func Request_Post_Login(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
	fmt.Println("got POST request to /login")

	var body LoginRequest
	req.ParseBodyJSON(&body)
	token, err := jwttokens.CreateToken()
	if err != nil {
		res.SendString("ERROR")
		return
	}
	resVal := &LoginResponse{
		Test: token,
	}
	res.SendJSON(resVal)
}
