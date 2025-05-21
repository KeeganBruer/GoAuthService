package controller_tests

import (
	controllers "go-auth-service/controllers"
	"testing"
)

var app *controllers.App

func TestMain(m *testing.M) {
	app = controllers.CreateApp()
	m.Run()

}
