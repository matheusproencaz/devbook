package routes

import (
	"net/http"
	"web/src/controllers"
)

var loginRoutes = []Route{
	{
		URI:       "/",
		Method:    http.MethodGet,
		Handler:   controllers.LoadLoginScreen,
		NeedsAuth: false,
	},
	{
		URI:       "/login",
		Method:    http.MethodGet,
		Handler:   controllers.LoadLoginScreen,
		NeedsAuth: false,
	},
	{
		URI:       "/login",
		Method:    http.MethodPost,
		Handler:   controllers.Login,
		NeedsAuth: false,
	},
}
