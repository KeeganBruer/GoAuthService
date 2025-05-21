package main

import (
	"fmt"
	main_app "go-auth-service/app"
)

func main() {
	PORT := 8080
	app := main_app.CreateApp()

	err := app.Listen(PORT, func(port int) {
		msg := fmt.Sprintf("Listening on port: %d", port)
		fmt.Println(msg)
	})

	if err != nil {
		panic(err)
	}
}
