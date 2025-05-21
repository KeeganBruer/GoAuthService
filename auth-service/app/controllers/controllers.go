package controllers

import (
	"go-auth-service/app/models"
	"kbrouter"
)

type Controller struct {
	Path   string
	Router *kbrouter.Router
	Models *models.Models
}

func (controller *Controller) GetRouter() *kbrouter.Router {
	return controller.Router
}
func (controller *Controller) AttachToRouter(baseRouter *kbrouter.Router) {
	baseRouter.AddSubRouter(controller.Path, controller.Router)
}
