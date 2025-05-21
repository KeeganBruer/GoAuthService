package controllers

import (
	"go-auth-service/app/configs"
	"go-auth-service/app/models"
	"kbrouter"
)

type Controller struct {
	Path       string
	Router     *kbrouter.Router
	Models     *models.Models
	AppConfigs *configs.AppConfigs
}

func (controller *Controller) GetRouter() *kbrouter.Router {
	return controller.Router
}
func (controller *Controller) AttachToRouter(baseRouter *kbrouter.Router) {
	baseRouter.AddSubRouter(controller.Path, controller.Router)
}
