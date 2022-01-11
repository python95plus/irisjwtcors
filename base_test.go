package main

import (
	"flag"
	"os"
	"testing"

	"github.com/iris-contrib/httpexpect/v2"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

const (
	baseUrl  = "/v1/admin/"
	loginUrl = baseUrl + "/login"
)

var (
	app *iris.Application
)

func TestMain(m *testing.M) {
	app = NewApp()
	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func Login(t *testing.T, object interface{}, StatusCode int, Status bool, Msg string) (e *httpexpect.Expect) {
	e = httptest.New(t, app, httptest.Configuration{Debug: true})
	e.POST(loginUrl).WithJSON(object).Expect().Status(StatusCode).JSON().Object().Values().Contains(Status, Msg)
	return
}
