package routes

import (
	"net/http"
	"web/src/controllers"
)

var userRoutes = []Route{
	{
		URI:       "/signup",
		Method:    http.MethodGet,
		Handler:   controllers.LoadSignupScreen,
		NeedsAuth: false,
	},
	{
		URI:       "/users",
		Method:    http.MethodPost,
		Handler:   controllers.CreateUser,
		NeedsAuth: false,
	},
}
