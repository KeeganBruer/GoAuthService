package app

import (
	"fmt"
)

func main() {
	port := 8080
	app := CreateApp()

	err := app.Listen(port, func() {
		msg := fmt.Sprintf("Listening on port: %d", port)
		fmt.Println(msg)
	})

	if err != nil {
		panic(err)
	}
}
