package app

import (
	controller_login "go-auth-service/controllers/login"
	controller_signup "go-auth-service/controllers/signup"
	"go-auth-service/models"
	"kbrouter"
)

func CreateApp() *kbrouter.Router {
	router := kbrouter.NewRouter()
	models.ConnectDB()

	//declare endpoints
	router.AddRoute("POST", "/login", controller_login.Login_PostRequest)
	router.AddRoute("POST", "/signup", controller_signup.Signup_PostRequest)

	return router
}
