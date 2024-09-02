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
}
