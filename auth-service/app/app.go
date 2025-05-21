package main_app

import (
	"go-auth-service/app/configs"
	"go-auth-service/app/controllers"
	auth_controller "go-auth-service/app/controllers/auth"
	session_controller "go-auth-service/app/controllers/session"
	swagger_controller "go-auth-service/app/controllers/swagger"
	token_controller "go-auth-service/app/controllers/token"
	"go-auth-service/app/models"
	"kbrouter"
)

type App struct {
	PublicRouter  *kbrouter.Router
	PrivateRouter *kbrouter.Router
}

func CreateApp() *App {
	CONFIGS := configs.NewAppConfigs()
	//Setup database connection
	models := models.ConnectDB(CONFIGS.Database)

	// Setup public API
	publicRouter := kbrouter.NewRouter()
	publicRouter.AddHealthRoute("/healthz")
	AttachControllers(publicRouter,
		swagger_controller.InitController(CONFIGS, models, false),
		auth_controller.InitController(CONFIGS, models),
		token_controller.InitController(CONFIGS, models),
	)

	// Setup internal API
	privateRouter := kbrouter.NewRouter()
	privateRouter.AddHealthRoute("/healthz")
	AttachControllers(privateRouter,
		swagger_controller.InitController(CONFIGS, models, true),
		session_controller.InitController(CONFIGS, models),
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
