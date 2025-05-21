package main

import (
	"fmt"
	main_app "go-auth-service/app"
	"go-auth-service/app/configs"
)

func main() {
	CONFIGS := configs.NewAppConfigs()

	app := main_app.CreateApp(CONFIGS)

	err := app.Listen(CONFIGS.InternalPorts, func(port int) {
		var msg string
		if port == CONFIGS.InternalPorts.PublicPort {
			msg = fmt.Sprintf("Listening on port: %d", CONFIGS.ExternalPorts.PublicPort)
		} else {
			msg = fmt.Sprintf("Listening on port: %d", CONFIGS.ExternalPorts.PrivatePort)
		}
		fmt.Println(msg)
	})

	if err != nil {
		panic(err)
	}
}
