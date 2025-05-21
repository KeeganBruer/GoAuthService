package token_controller

import (
	"go-auth-service/app/configs"
	"go-auth-service/app/controllers"
	"go-auth-service/app/models"
	"kbrouter"
)

type TokenController struct {
	controllers.Controller
}

func InitController(CONFIGS *configs.AppConfigs, models *models.Models) *controllers.Controller {
	router := kbrouter.NewRouter()
	controller := &TokenController{
		controllers.Controller{
			Path:       "/token",
			Models:     models,
			Router:     router,
			AppConfigs: CONFIGS,
		},
	}

	router.AddRoute("GET", "/verify", controller.Verify_GetRequest)
	router.AddRoute("POST", "/refresh", controller.Refresh_PostRequest)

	return &controller.Controller
}
