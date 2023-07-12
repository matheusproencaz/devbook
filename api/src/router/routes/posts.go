package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesPosts = []Route{
	{
		URI:            "/posts",
		Method:         http.MethodPost,
		Function:       controllers.CreatePost,
		Authentication: true,
	},
	{
		URI:            "/posts",
		Method:         http.MethodGet,
		Function:       controllers.GetPosts,
		Authentication: true,
	},
	{
		URI:            "/posts/{postId}",
		Method:         http.MethodGet,
		Function:       controllers.GetPostByID,
		Authentication: true,
	},
	{
		URI:            "/posts/{postId}",
		Method:         http.MethodPut,
		Function:       controllers.UpdatePost,
		Authentication: true,
	},
	{
		URI:            "/posts/{postId}",
		Method:         http.MethodDelete,
		Function:       controllers.DeletePost,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}/posts",
		Method:         http.MethodGet,
		Function:       controllers.GetPostByUser,
		Authentication: true,
	},
	{
		URI:            "/posts/{postId}/like",
		Method:         http.MethodPost,
		Function:       controllers.LikePost,
		Authentication: true,
	},
	{
		URI:            "/posts/{postId}/unlike",
		Method:         http.MethodPost,
		Function:       controllers.UnLikePost,
		Authentication: true,
	},
}
