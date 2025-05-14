package app

import (
	controller_login "go-auth-service/controllers/login"
	controller_signup "go-auth-service/controllers/signup"
	"kbrouter"
)

func CreateApp() *kbrouter.Router {
	router := kbrouter.NewRouter()

	//declare endpoints
	router.AddRoute("POST", "/login", controller_login.Login_PostRequest)
	router.AddRoute("POST", "/signup", controller_signup.Signup_PostRequest)

	return router
}
