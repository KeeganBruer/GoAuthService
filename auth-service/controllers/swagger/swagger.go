package controller_swagger

import (
	"fmt"
	"kbrouter"
	"path/filepath"
)

func CreateSwaggerRouter() *kbrouter.Router {
	router := kbrouter.NewRouter()

	//serve SwaggerUI
	router.AddRoute("GET", "/", func(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
		absPath, err := filepath.Abs("../swagger/index.html")
		if err != nil {
			res.SetStatusCode(400).SendString(fmt.Sprintf("Error getting absolute path: %v", err))
			return
		}
		res.SetHeader("Content-Type", "text/html")
		res.SendFile(absPath)
	})
	//Server swagger openapi yaml
	router.AddRoute("GET", "/swagger.yaml", func(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
		absPath, err := filepath.Abs("../swagger/swagger.yaml")
		if err != nil {
			res.SetStatusCode(400).SendString(fmt.Sprintf("Error getting absolute path: %v", err))
			return
		}
		res.SetHeader("Content-Type", "application/x-yaml")
		res.SendFile(absPath)
	})
	return router
}
