package main

import (
	"fmt"
	"go-auth-service/controllers"
)

func main() {
	port := 8080
	app := controllers.CreateApp()

	err := app.Listen(port, func(port int) {
		msg := fmt.Sprintf("Listening on port: %d", port)
		fmt.Println(msg)
	})

	if err != nil {
		panic(err)
	}
}
