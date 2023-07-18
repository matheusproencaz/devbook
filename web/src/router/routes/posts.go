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
	{
		URI:       "/posts/{postId}/like",
		Method:    http.MethodPost,
		Handler:   controllers.LikePost,
		NeedsAuth: true,
	},
	{
		URI:       "/posts/{postId}/dislike",
		Method:    http.MethodPost,
		Handler:   controllers.DislikePost,
		NeedsAuth: true,
	},
	{
		URI:       "/posts/{postId}/edit",
		Method:    http.MethodGet,
		Handler:   controllers.LoadEditionPostPage,
		NeedsAuth: true,
	},
	{
		URI:       "/posts/{postId}",
		Method:    http.MethodPut,
		Handler:   controllers.UpdatePost,
		NeedsAuth: true,
	},
	{
		URI:       "/posts/{postId}",
		Method:    http.MethodDelete,
		Handler:   controllers.DeletePost,
		NeedsAuth: true,
	},
}
