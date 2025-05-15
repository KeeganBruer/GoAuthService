package controller_tests

import (
	controllers "go-auth-service/controllers"
	"kbrouter"
	"testing"
)

var router *kbrouter.Router

func TestMain(m *testing.M) {
	router = controllers.CreateApp()
	m.Run()

}
