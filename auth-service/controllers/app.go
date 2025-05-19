package controllers

import (
	"fmt"
	controller_api_key "go-auth-service/controllers/api_key"
	controller_login "go-auth-service/controllers/login"
	controller_signup "go-auth-service/controllers/signup"
	controller_token "go-auth-service/controllers/token"
	"go-auth-service/models"
	"kbrouter"
	"path/filepath"
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
	privateRouter.AddRoute("GET", "/swagger", func(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
		absPath, err := filepath.Abs("../swagger/index.html")
		if err != nil {
			res.SetStatusCode(400).SendString(fmt.Sprintf("Error getting absolute path: %v", err))
			return
		}
		res.SetHeader("Content-Type", "text/html")
		res.SendFile(absPath)
	})
	privateRouter.AddRoute("GET", "/swagger/swagger.yaml", func(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
		absPath, err := filepath.Abs("../swagger/swagger.yaml")
		if err != nil {
			res.SetStatusCode(400).SendString(fmt.Sprintf("Error getting absolute path: %v", err))
			return
		}
		res.SetHeader("Content-Type", "application/x-yaml")
		res.SendFile(absPath)
	})

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
