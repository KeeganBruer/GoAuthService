package main_app

import (
	auth_controller "go-auth-service/app/controllers/auth"
	swagger_controller "go-auth-service/app/controllers/swagger"
	"go-auth-service/app/models"
	"kbrouter"
)

type App struct {
	PublicRouter  *kbrouter.Router
	PrivateRouter *kbrouter.Router
}

func CreateApp() *App {

	//Setup database connection
	models := models.ConnectDB()

	// Setup public API
	publicRouter := kbrouter.NewRouter()
	publicRouter.AddHealthRoute("/healthz")
	auth_controller.InitController(models).AttachToRouter(publicRouter, "/auth")
	swagger_controller.InitController(models, false).AttachToRouter(publicRouter, "/swagger")

	privateRouter := kbrouter.NewRouter()
	privateRouter.AddHealthRoute("/healthz")
	swagger_controller.InitController(models, true).AttachToRouter(privateRouter, "/swagger")

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
