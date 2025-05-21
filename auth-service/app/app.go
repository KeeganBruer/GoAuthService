package main_app

import (
	"go-auth-service/app/controllers"
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
	AttachControllers(publicRouter,
		auth_controller.InitController(models),
		swagger_controller.InitController(models, false),
	)

	privateRouter := kbrouter.NewRouter()
	privateRouter.AddHealthRoute("/healthz")
	AttachControllers(privateRouter,
		swagger_controller.InitController(models, true),
	)

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
func AttachControllers(router *kbrouter.Router, controllers ...*controllers.Controller) {
	for i := range controllers {
		controllers[i].AttachToRouter(router)
	}
}
