package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.CreateUser,
		Authentication: false,
	},
	{
		URI:            "/users",
		Method:         http.MethodGet,
		Function:       controllers.GetUsers,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodGet,
		Function:       controllers.GetUserById,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodPut,
		Function:       controllers.UpdateUser,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodDelete,
		Function:       controllers.DeleteUser,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}/follow",
		Method:         http.MethodPost,
		Function:       controllers.FollowUser,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}/unfollow",
		Method:         http.MethodPost,
		Function:       controllers.UnFollowUser,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}/followers",
		Method:         http.MethodGet,
		Function:       controllers.GetFollowers,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}/following",
		Method:         http.MethodGet,
		Function:       controllers.GetFollowing,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}/updatepassword",
		Method:         http.MethodPost,
		Function:       controllers.UpdatePassword,
		Authentication: true,
	},
}
