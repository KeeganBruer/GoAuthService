package main

import (
	"fmt"
	"go-auth-service/routes"
	"kbrouter"
)

func main() {
	port := 8080

	router := kbrouter.NewRouter()
	router.AddRoute("POST", "/login", routes.Request_Post_Login)

	err := router.Listen(port, func() {
		msg := fmt.Sprintf("Listening on port: %d", port)
		fmt.Println(msg)
	})

	if err != nil {
		panic(err)
	}
}
