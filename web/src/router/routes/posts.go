package routes

import (
	"net/http"
	"web/src/controllers"
)

var postsRoutes = []Route{
	{
		URI:       "/posts",
		Method:    http.MethodPost,
		Handler:   controllers.CreatePost,
		NeedsAuth: true,
	},
}
