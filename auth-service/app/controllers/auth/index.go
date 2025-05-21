package auth_controller

import (
	"go-auth-service/app/configs"
	"go-auth-service/app/controllers"
	"go-auth-service/app/models"
	"kbrouter"
)

type AuthController struct {
	controllers.Controller
}

func InitController(CONFIGS *configs.AppConfigs, models *models.Models) *controllers.Controller {
	router := kbrouter.NewRouter()
	controller := &AuthController{
		controllers.Controller{
			Path:       "/auth",
			Models:     models,
			Router:     router,
			AppConfigs: CONFIGS,
		},
	}

	router.AddRoute("POST", "/login", controller.Login_PostRequest)
	router.AddRoute("POST", "/signup", controller.Signup_PostRequest)

	return &controller.Controller
}
