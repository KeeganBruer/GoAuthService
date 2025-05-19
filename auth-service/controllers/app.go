package controllers

import (
	controller_api_key "go-auth-service/controllers/api_key"
	controller_login "go-auth-service/controllers/login"
	controller_signup "go-auth-service/controllers/signup"
	controller_token "go-auth-service/controllers/token"
	"go-auth-service/models"
	"kbrouter"
)

type App struct {
	PublicRouter  *kbrouter.Router
	PrivateRouter *kbrouter.Router
}

func CreateApp() *App {

	//Setup database connection
	models.ConnectDB()

	// Setup  API router
	publicRouter := kbrouter.NewRouter()
	//declare endpoints
	publicRouter.AddRoute("POST", "/login", controller_login.Login_PostRequest)
	publicRouter.AddRoute("POST", "/signup", controller_signup.Signup_PostRequest)
	publicRouter.AddRoute("GET", "/token/verify", controller_token.Verify_GetRequest)
	publicRouter.AddRoute("POST", "/token/refresh", controller_token.Refresh_PostRequest)
	publicRouter.AddRoute("GET", "/api_key/verify", controller_api_key.Verify_GetRequest)

	privateRouter := kbrouter.NewRouter()

	app := &App{
		PublicRouter:  publicRouter,
		PrivateRouter: privateRouter,
	}
	return app
}

func (app *App) Listen(port int, cb func(port int)) error {
	go app.PrivateRouter.Listen(port+1, cb)
	app.PublicRouter.Listen(port, cb)
	return nil
}
