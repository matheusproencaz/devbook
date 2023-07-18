package routes

import (
	"net/http"
	"web/src/controllers"
)

var userRoutes = []Route{
	{
		URI:       "/signup",
		Method:    http.MethodGet,
		Handler:   controllers.LoadSignupPage,
		NeedsAuth: false,
	},
	{
		URI:       "/users",
		Method:    http.MethodPost,
		Handler:   controllers.CreateUser,
		NeedsAuth: false,
	},
	{
		URI:       "/search-users",
		Method:    http.MethodGet,
		Handler:   controllers.LoadUsersPage,
		NeedsAuth: true,
	},
	{
		URI:       "/users/{userId}",
		Method:    http.MethodGet,
		Handler:   controllers.LoadUserProfilePage,
		NeedsAuth: true,
	},
	{
		URI:       "/users/{userId}/unfollow",
		Method:    http.MethodPost,
		Handler:   controllers.Unfollow,
		NeedsAuth: true,
	},
	{
		URI:       "/users/{userId}/follow",
		Method:    http.MethodPost,
		Handler:   controllers.Follow,
		NeedsAuth: true,
	},
	{
		URI:       "/profile",
		Method:    http.MethodGet,
		Handler:   controllers.LoadLoggedInUserProfilePage,
		NeedsAuth: true,
	},
	{
		URI:       "/edit-user",
		Method:    http.MethodGet,
		Handler:   controllers.LoadUserEditionPage,
		NeedsAuth: true,
	},
	{
		URI:       "/edit-user",
		Method:    http.MethodPut,
		Handler:   controllers.EditUser,
		NeedsAuth: true,
	},
	{
		URI:       "/edit-password",
		Method:    http.MethodGet,
		Handler:   controllers.LoadUserPasswordEditionPage,
		NeedsAuth: true,
	},
	{
		URI:       "/edit-password",
		Method:    http.MethodPost,
		Handler:   controllers.UpdatePassword,
		NeedsAuth: true,
	},
	{
		URI:       "/delete-user",
		Method:    http.MethodDelete,
		Handler:   controllers.DeleteUser,
		NeedsAuth: true,
	},
}
