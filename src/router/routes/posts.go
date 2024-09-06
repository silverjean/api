package routes

import (
	"api/src/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		Uri:    	 "/posts",
		Method: 	 http.MethodPost,
		Function: 	 controllers.CreatePost,
		RequireAuth: true,
	},
	{
		Uri:    	 "/posts",
		Method: 	 http.MethodGet,
		Function: 	 controllers.FindPosts,
		RequireAuth: true,
	},
	{
		Uri:    	 "/posts/{postId}",
		Method: 	 http.MethodGet,
		Function: 	 controllers.FindPost,
		RequireAuth: true,
	},
	{
		Uri:    	 "/posts/{postId}",
		Method: 	 http.MethodPut,
		Function: 	 controllers.UpdatePost,
		RequireAuth: true,
	},
	{
		Uri:    	 "/posts/{postId}",
		Method: 	 http.MethodDelete,
		Function: 	 controllers.DeletePost,
		RequireAuth: true,
	},
}