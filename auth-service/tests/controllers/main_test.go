package controller_tests

import (
	app "go-auth-service"
	"kbrouter"
	"testing"
)

var router *kbrouter.Router

func TestMain(m *testing.M) {
	router = app.CreateApp()
	m.Run()

}
