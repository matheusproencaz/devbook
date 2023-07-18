package routes

import (
	"net/http"
	"web/src/controllers"
)

var logoutRoute = Route{
	URI:       "/logout",
	Method:    http.MethodGet,
	Handler:   controllers.Logout,
	NeedsAuth: true,
}
