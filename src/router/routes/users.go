package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		Uri:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.FindUsers,
		RequireAuth: true,
	},
	{
		Uri:         "/users/{userId}",
		Method:      http.MethodGet,
		Function:    controllers.FindUser,
		RequireAuth: true,
	},
	{
		Uri:         "/users/{userId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		Uri:         "/users/{userId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUsers,
		RequireAuth: true,
	},
	{
		Uri:         "/users/{userId}/follow",
		Method:      http.MethodPost,
		Function:    controllers.FollowUser,
		RequireAuth: true,
	},
	{
		Uri:         "/users/{userId}/unfollow",
		Method:      http.MethodPost,
		Function:    controllers.UnfollowUser,
		RequireAuth: true,
	},
	{
		Uri:         "/users/{userId}/followers",
		Method:      http.MethodGet,
		Function:    controllers.FollowersUser,
		RequireAuth: true,
	},
	{
		Uri:         "/users/{userId}/following",
		Method:      http.MethodGet,
		Function:    controllers.FollowingUser,
		RequireAuth: true,
	},
	{
		Uri:         "/users/{userId}/update-password",
		Method:      http.MethodPost,
		Function:    controllers.UpdatePassUser,
		RequireAuth: true,
	},
}
