package controller_token

import (
	"go-auth-service/app/controllers"
	"go-auth-service/app/models"
	"kbrouter"
)

type TokenController struct {
	controllers.Controller
}

func InitController(models *models.Models) *controllers.Controller {
	router := kbrouter.NewRouter()
	controller := &TokenController{
		controllers.Controller{
			Path:   "/token",
			Models: models,
			Router: router,
		},
	}

	router.AddRoute("GET", "/verify", controller.Verify_GetRequest)
	router.AddRoute("POST", "/refresh", controller.Refresh_PostRequest)

	return &controller.Controller
}
