package controllers

import (
	controller_api_key "go-auth-service/controllers/api_key"
	controller_login "go-auth-service/controllers/login"
	controller_signup "go-auth-service/controllers/signup"
	controller_swagger "go-auth-service/controllers/swagger"
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

	// Setup public API
	publicRouter := kbrouter.NewRouter()
	publicRouter.AddHealthRoute("/healthz")

	publicRouter.AddRoute("POST", "/login", controller_login.Login_PostRequest)
	publicRouter.AddRoute("POST", "/signup", controller_signup.Signup_PostRequest)

	tokenRouter := controller_token.CreateTokenRouter()
	publicRouter.AddSubRouter("/token", tokenRouter)

	publicRouter.AddRoute("GET", "/api_key/verify", controller_api_key.Verify_GetRequest)

	// Private service for internal communication
	privateRouter := kbrouter.NewRouter()
	privateRouter.AddHealthRoute("/healthz")

	swagger := controller_swagger.CreateSwaggerRouter()
	privateRouter.AddSubRouter("/swagger", swagger)

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
