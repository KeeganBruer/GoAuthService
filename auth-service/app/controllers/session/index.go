package session_controller

import (
	"fmt"
	"go-auth-service/app/configs"
	"go-auth-service/app/controllers"
	"go-auth-service/app/models"
	"kbrouter"
)

type SessionController struct {
	controllers.Controller
}

func InitController(CONFIGS *configs.AppConfigs, models *models.Models) *controllers.Controller {
	router := kbrouter.NewRouter()
	controller := &SessionController{
		controllers.Controller{
			Path:       "/session",
			Models:     models,
			Router:     router,
			AppConfigs: CONFIGS,
		},
	}

	router.AddRoute("GET", "/$sessionID", func(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
		sessionID, err := req.GetIntParam("sessionID")
		if err != nil {
			res.SetStatusCode(400).SendString("Request did not have a valid sessionID")
			return
		}
		fmt.Println("Session ID:", sessionID)
		res.SendJSONString(
			`{
				"id":%d,
				"user": {
					"id":"ahh",
					"username":"keegan"
				}
			}`,
			sessionID,
		)
	})

	return &controller.Controller
}
