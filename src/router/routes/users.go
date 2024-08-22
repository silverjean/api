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
		RequestAuth: false,
	},
	{
		Uri:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.FindUsers,
		RequestAuth: false,
	},
	{
		Uri:         "/users/{userId}",
		Method:      http.MethodGet,
		Function:    controllers.FindUser,
		RequestAuth: false,
	},
	{
		Uri:         "/users/{userId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequestAuth: false,
	},
	{
		Uri:         "/users/{userId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUsers,
		RequestAuth: false,
	},
}
