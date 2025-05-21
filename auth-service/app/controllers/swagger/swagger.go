package swagger_controller

import (
	"fmt"
	"go-auth-service/app/configs"
	"go-auth-service/app/controllers"
	"go-auth-service/app/middleware"
	"go-auth-service/app/models"
	"kbrouter"
	"os"
	"path/filepath"
	"strings"
)

type SwaggerController struct {
	controllers.Controller
}

func InitController(CONFIGS *configs.AppConfigs, models *models.Models, isPrivate bool) *controllers.Controller {
	router := kbrouter.NewRouter()
	controller := &SwaggerController{
		controllers.Controller{
			Path:       "/swagger",
			Models:     models,
			Router:     router,
			AppConfigs: CONFIGS,
		},
	}
	if isPrivate && CONFIGS.IsDev {
		fmt.Printf(
			"\x1b]8;;http://localhost:%d/swagger\x07%s\x1b]8;;\x07\u001b[0m\n",
			CONFIGS.ExternalPorts.PrivatePort,
			"Open Swagger Documentation",
		)
	} else if !isPrivate && !CONFIGS.IsDev {
		fmt.Printf(
			"\x1b]8;;http://localhost:%d/swagger\x07%s\x1b]8;;\x07\u001b[0m\n",
			CONFIGS.ExternalPorts.PublicPort,
			"Open Swagger Documentation",
		)
	}
	//
	router.AddRoute("GET", "/", SendSwaggerUI("/swagger-public.yaml", false))
	router.AddRoute("GET", "/swagger-public.yaml", middleware.ServeStaticFile("../swagger/swagger-public.yaml"))
	if isPrivate {
		router.AddRoute("GET", "/", middleware.ServeStaticFile("../swagger/index.html"))
		router.AddRoute("GET", "/public", SendSwaggerUI("/swagger-public.yaml", true))
		router.AddRoute("GET", "/private", SendSwaggerUI("/swagger-private.yaml", true))
		router.AddRoute("GET", "/swagger-private.yaml", middleware.ServeStaticFile("../swagger/swagger-private.yaml"))
	}

	return &controller.Controller
}

func SendSwaggerUI(yamlFile string, backButton bool) kbrouter.KBRouteHandler {
	return func(req *kbrouter.KBRequest, res *kbrouter.KBResponse) {
		absPath, err := filepath.Abs("../swagger/swagger.html")
		if err != nil {
			res.SetStatusCode(400).SendString(fmt.Sprintf("Error getting absolute path: %v", err))
			return
		}
		data, err := os.ReadFile(absPath)
		if err != nil {
			res.SetStatusCode(400).SendString(fmt.Sprintf("Error reading file: %v", err))
			return
		}
		fileContent := string(data)
		fileContent = strings.Replace(fileContent, "/swagger.yaml", yamlFile, 1)
		if backButton {
			fileContent = strings.Replace(fileContent, `<a id="back-button" style="display:none;"`, `<a id="back-button"`, 1)
		}
		if err != nil {
			res.SetStatusCode(400).SendString(fmt.Sprintf("Error getting file content: %v", err))
			return
		}
		res.SetHeader("Content-Type", "text/html")
		res.SendString(fileContent)
	}
}
