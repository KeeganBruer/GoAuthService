package controllers

import (
	controller_login "go-auth-service/controllers/login"
	controller_signup "go-auth-service/controllers/signup"
	controller_token "go-auth-service/controllers/token"
	"go-auth-service/models"
	"kbrouter"
)

func CreateApp() *kbrouter.Router {

	//Setup database connection
	models.ConnectDB()

	// Setup  API router
	router := kbrouter.NewRouter()
	//declare endpoints
	router.AddRoute("POST", "/login", controller_login.Login_PostRequest)
	router.AddRoute("POST", "/signup", controller_signup.Signup_PostRequest)
	router.AddRoute("GET", "/token/verify", controller_token.Verify_GetRequest)
	router.AddRoute("POST", "/token/refresh", controller_token.Refresh_PostRequest)

	return router
}
