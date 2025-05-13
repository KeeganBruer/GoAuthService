package app

import (
	"fmt"
	"go-auth-service/controllers"
	"kbrouter"
)

func CreateApp() *kbrouter.Router {
	router := kbrouter.NewRouter()
	router.AddRoute("POST", "/login", controllers.Request_Post_Login)
	router.AddRoute("GET", "/wild", func(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
		// Implementation for creating a new product
		fmt.Println("got request to /wild")
		res.SendString("OKAY\n")
	})
	return router
}
