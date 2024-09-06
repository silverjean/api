package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Routes represents all routes from api
type Route struct {
	Uri      	string
	Method   	string
	Function 	func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}
// Configure add all routes insider the router
func Configure(router *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, postRoutes...)

	for _, route := range routes {
		if route.RequireAuth {
			router.HandleFunc(route.Uri, middlewares.Logger(middlewares.Auth(route.Function))).Methods(route.Method)
		}
		router.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)
	}

	return router
}