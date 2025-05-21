package main_app

import (
	"go-auth-service/app/configs"
	"go-auth-service/app/controllers"
	auth_controller "go-auth-service/app/controllers/auth"
	session_controller "go-auth-service/app/controllers/session"
	swagger_controller "go-auth-service/app/controllers/swagger"
	token_controller "go-auth-service/app/controllers/token"
	"go-auth-service/app/models"
	timedjobs "go-auth-service/app/timed_jobs"
	intervaljobs "intervaljobs"
	"kbrouter"
	"time"
)

type App struct {
	PublicRouter  *kbrouter.Router
	PrivateRouter *kbrouter.Router
	IntervalJobs  *intervaljobs.IntervalJobManager
}

func CreateApp(CONFIGS *configs.AppConfigs) *App {
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

	IntervalJobs := intervaljobs.NewIntervalJobManager(intervaljobs.ManagerConfigs{
		Interval: 30 * time.Minute,
	})
	IntervalJobs.AddJob(intervaljobs.JobConfig{
		Interval: &intervaljobs.IntervalJobTime{
			Hr: 1,
		},
		Offset: &intervaljobs.IntervalJobTime{
			Min: 5,
		},
		Handler: timedjobs.CleanSessions,
	})

	app := &App{
		IntervalJobs:  IntervalJobs,
		PublicRouter:  publicRouter,
		PrivateRouter: privateRouter,
	}
	return app
}

func (app *App) Listen(ports configs.Ports, cb func(port int)) error {
	go app.IntervalJobs.Start()
	go app.PrivateRouter.Listen(ports.PrivatePort, cb)
	app.PublicRouter.Listen(ports.PublicPort, cb)
	return nil
}
func AttachControllers(router *kbrouter.Router, controllers ...*controllers.Controller) {
	for i := range controllers {
		controllers[i].AttachToRouter(router)
	}
}
