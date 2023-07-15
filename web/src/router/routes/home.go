package routes

import (
	"net/http"
	"web/src/controllers"
)

var homeRoute = Route{
	URI:       "/home",
	Method:    http.MethodGet,
	Handler:   controllers.LoadHomeScreen,
	NeedsAuth: true,
}
