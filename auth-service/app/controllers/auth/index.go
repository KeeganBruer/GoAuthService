package auth_controller

import (
	"go-auth-service/app/controllers"
	"go-auth-service/app/models"
	"kbrouter"
)

type AuthController struct {
	controllers.Controller
}

func InitController(models *models.Models) *AuthController {
	router := kbrouter.NewRouter()
	controller := &AuthController{
		controllers.Controller{
			Models: models,
			Router: router,
		},
	}

	router.AddRoute("POST", "/login", controller.Login_PostRequest)
	router.AddRoute("POST", "/signup", controller.Signup_PostRequest)

	return controller
}
