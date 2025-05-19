package controller_token

import (
	"kbrouter"
)

func CreateTokenRouter() *kbrouter.Router {
	router := kbrouter.NewRouter()
	router.AddRoute("GET", "/verify", Verify_GetRequest)
	router.AddRoute("POST", "/refresh", Refresh_PostRequest)
	return router
}
