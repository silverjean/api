package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes represents all routes from api
type Route struct {
	Uri      	string
	Method   	string
	Function 	func(http.ResponseWriter, *http.Request)
	RequestAuth bool
}
// Configure add all routes insider the router
func Configure(router *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		router.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}

	return router
}